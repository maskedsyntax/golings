package cli

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"

	"github.com/maskedsyntax/golings/internal/app"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/exitcode"
	"github.com/maskedsyntax/golings/internal/render"
	"github.com/maskedsyntax/golings/internal/toolchain"
	"github.com/maskedsyntax/golings/internal/watcher"
)

func cmdWatch() *command {
	var autoAdvance bool
	return &command{
		name:    "watch",
		summary: "watch the current exercise and re-check it on every save (the default command)",
		setup: func(fs *flag.FlagSet) {
			fs.BoolVar(&autoAdvance, "auto-advance", false, "move to the next exercise automatically after a pass")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) > 0 {
				return exitcode.New(exitcode.Usage, "watch takes no arguments; to check a specific exercise use `golings verify <id>`")
			}
			if e.g.json {
				return exitcode.New(exitcode.Usage, "watch is interactive; use `golings verify --json` for machine-readable results")
			}
			a, err := e.app()
			if err != nil {
				return err
			}
			tc, err := e.toolchain()
			if err != nil {
				return err
			}
			return newWatchUI(e, a, tc, autoAdvance).run()
		},
	}
}

// key is one learner command, from a keypress (TTY) or an input line.
type key byte

const (
	keyHint     key = 'h'
	keyNext     key = 'n'
	keyPrev     key = 'p'
	keyList     key = 'l'
	keyRerun    key = 'r'
	keyQuit     key = 'q'
	keyHelp     key = '?'
	keyInterupt key = 3 // Ctrl+C in raw mode
	keyEOF      key = 4 // Ctrl+D
)

type runOutcome struct {
	gen int
	ev  *app.Evaluation
	err error
}

type watchUI struct {
	e           *env
	a           *app.App
	tc          *toolchain.Info
	tty         bool
	out         io.Writer
	autoAdvance bool

	ex        *catalog.Exercise
	gen       int
	cancelRun context.CancelFunc
	results   chan runOutcome
	changes   <-chan struct{}
	stopWatch context.CancelFunc

	// screen state (TTY mode)
	checking bool
	result   *app.Evaluation
	runErr   error
	notice   string
}

func newWatchUI(e *env, a *app.App, tc *toolchain.Info, autoAdvance bool) *watchUI {
	return &watchUI{e: e, a: a, tc: tc, out: e.stdout, autoAdvance: autoAdvance, results: make(chan runOutcome, 4)}
}

func (w *watchUI) run() (err error) {
	ctx := w.e.ctx
	keys := make(chan key, 8)

	if w.e.interactive() {
		in := w.e.stdin.(*os.File)
		state, rawErr := term.MakeRaw(int(in.Fd()))
		if rawErr == nil {
			w.tty = true
			// Raw mode disables output post-processing: translate \n.
			w.out = crlfWriter{w.e.stdout}
			defer func() {
				term.Restore(int(in.Fd()), state)
				fmt.Fprintln(w.e.stdout)
			}()
		}
	}
	go readKeys(w.e.stdin, w.tty, keys)

	st, err := w.a.Store.Load()
	if err != nil {
		return err
	}
	ex := w.a.Current(st)
	if ex == nil {
		fmt.Fprintln(w.out, w.e.style.Green("🎉 Every exercise is complete. Well done!"))
		return nil
	}
	if err := w.switchTo(ctx, ex); err != nil {
		return err
	}
	defer func() {
		if w.cancelRun != nil {
			w.cancelRun()
		}
		if w.stopWatch != nil {
			w.stopWatch()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case _, ok := <-w.changes:
			if !ok {
				w.changes = nil
				continue
			}
			w.startRun(ctx, "change detected")

		case r := <-w.results:
			if r.gen != w.gen {
				continue // a newer save superseded this run
			}
			w.checking, w.result, w.runErr = false, r.ev, r.err
			w.notice = ""
			w.showResult()
			if r.err == nil && r.ev.Result.Passed() && w.autoAdvance {
				if next := w.nextExercise(); next != nil && next.ID != w.ex.ID {
					if err := w.switchTo(ctx, next); err != nil {
						return err
					}
				}
			}

		case k := <-keys:
			done, err := w.handleKey(ctx, k)
			if done || err != nil {
				return err
			}
		}
	}
}

func (w *watchUI) handleKey(ctx context.Context, k key) (quit bool, err error) {
	switch k {
	case keyQuit, keyEOF:
		return true, nil
	case keyInterupt:
		return true, silentExit(exitcode.Interrupted)
	case keyRerun:
		w.startRun(ctx, "re-run requested")
	case keyHint:
		h, err := w.a.RevealHint(w.ex, 0)
		if err != nil {
			return false, err
		}
		w.redraw()
		if !w.tty {
			printHint(&env{stdout: w.out, style: w.e.style}, w.ex, h)
		}
	case keyNext, keyPrev:
		var target *catalog.Exercise
		if k == keyNext {
			target = w.nextExercise()
		} else if w.ex.Ordinal > 1 {
			target = w.a.Catalog.Exercises()[w.ex.Ordinal-2]
		}
		if target == nil {
			w.say("There is no exercise in that direction.")
			return false, nil
		}
		return false, w.switchTo(ctx, target)
	case keyList:
		w.say(w.listText())
	case keyHelp:
		w.say(helpText)
	}
	return false, nil
}

const helpText = "Keys: h hint · n next exercise · p previous · l list · r re-run · q quit.\n" +
	"Edit the exercise file in any editor; it is re-checked every time you save."

// switchTo makes ex the current exercise, watches its files and checks it.
func (w *watchUI) switchTo(ctx context.Context, ex *catalog.Exercise) error {
	if err := w.a.SelectCurrent(ex.ID); err != nil {
		return err
	}
	w.ex, w.result, w.runErr, w.notice = ex, nil, nil, ""
	if w.stopWatch != nil {
		w.stopWatch()
	}
	wctx, stop := context.WithCancel(ctx)
	w.stopWatch = stop
	paths := make([]string, len(ex.StarterFiles))
	for i, f := range ex.StarterFiles {
		paths[i] = filepath.Join(w.a.Workspace.ExerciseDir(ex), f)
	}
	w.changes = watcher.Watch(wctx, paths, watcher.Options{})
	if !w.tty {
		fmt.Fprintln(w.out)
		w.header()
	}
	w.startRun(ctx, "")
	return nil
}

// startRun cancels any in-flight evaluation and starts a new one. Only the
// newest generation's result is ever displayed.
func (w *watchUI) startRun(ctx context.Context, reason string) {
	if w.cancelRun != nil {
		w.cancelRun()
	}
	w.gen++
	gen, ex := w.gen, w.ex
	rctx, cancel := context.WithCancel(ctx)
	w.cancelRun = cancel
	w.checking = true
	if w.tty {
		w.redraw()
	} else {
		msg := "Checking " + ex.ID
		if reason != "" {
			msg += " (" + reason + ")"
		}
		fmt.Fprintf(w.out, "\n%s\n", w.e.style.Dim("— "+msg+" —"))
	}
	go func() {
		ev, err := w.a.Evaluate(rctx, w.tc, ex, w.e.g.timeout)
		if rctx.Err() != nil {
			return // superseded or shutting down
		}
		w.results <- runOutcome{gen: gen, ev: ev, err: err}
	}()
}

func (w *watchUI) nextExercise() *catalog.Exercise {
	st, err := w.a.Store.Load()
	if err != nil {
		return nil
	}
	return pickNext(w.a, st)
}

func (w *watchUI) header() {
	topic := w.ex.Topic
	if t := w.a.Catalog.Topic(w.ex.Topic); t != nil {
		topic = t.Title
	}
	render.ExerciseHeader(w.out, w.e.style, w.ex, len(w.a.Catalog.Exercises()), topic, w.a.Workspace.RelExerciseDir(w.ex))
}

// say shows a transient message.
func (w *watchUI) say(msg string) {
	if w.tty {
		w.notice = msg
		w.redraw()
		return
	}
	fmt.Fprintln(w.out, msg)
}

func (w *watchUI) showResult() {
	if w.tty {
		w.redraw()
		return
	}
	w.body(w.out)
}

// body writes the result section.
func (w *watchUI) body(out io.Writer) {
	s := w.e.style
	switch {
	case w.checking:
		fmt.Fprintln(out, s.Dim("Checking…"))
	case w.runErr != nil:
		fmt.Fprintf(out, "%s %v\n", s.Red("Error:"), w.runErr)
	case w.result != nil:
		render.Result(out, s, w.result.Result, w.e.g.verbose)
		if w.result.Result.Passed() {
			if w.autoAdvance {
				fmt.Fprintln(out, s.Green("Exercise complete! Moving on…"))
			} else {
				fmt.Fprintln(out, s.Green("Exercise complete! Press n for the next exercise."))
			}
		}
	}
}

// redraw repaints the whole screen (TTY mode only).
func (w *watchUI) redraw() {
	if !w.tty {
		return
	}
	var b bytes.Buffer
	b.WriteString("\x1b[H\x1b[2J") // home + clear
	s := w.e.style
	fmt.Fprintf(&b, "%s  %s\n\n", s.Bold("GOLINGS"), s.Dim("learn Go by fixing Go"))
	topic := w.ex.Topic
	if t := w.a.Catalog.Topic(w.ex.Topic); t != nil {
		topic = t.Title
	}
	render.ExerciseHeader(&b, s, w.ex, len(w.a.Catalog.Exercises()), topic, w.a.Workspace.RelExerciseDir(w.ex))
	b.WriteString("\n")
	w.body(&b)
	if st, err := w.a.Store.Load(); err == nil {
		hints := w.a.RevealedHints(w.ex, st)
		for i, h := range hints {
			fmt.Fprintf(&b, "\n%s %s", s.Yellow(fmt.Sprintf("Hint %d/%d:", i+1, len(w.ex.Hints))), h)
		}
		if len(hints) > 0 {
			b.WriteString("\n")
		}
	}
	if w.notice != "" {
		fmt.Fprintf(&b, "\n%s\n", w.notice)
	}
	fmt.Fprintf(&b, "\n%s\n", s.Dim("[h] hint  [n] next  [p] prev  [l] list  [r] re-run  [q] quit"))
	w.out.Write(b.Bytes())
}

func (w *watchUI) listText() string {
	st, err := w.a.Store.Load()
	if err != nil {
		return err.Error()
	}
	var b strings.Builder
	for _, v := range w.a.Views(st) {
		marker := " "
		if v.Exercise.ID == w.ex.ID {
			marker = ">"
		}
		fmt.Fprintf(&b, "%s %3d  %-18s %s\n", marker, v.Exercise.Ordinal, v.Exercise.ID, statusCell(w.e.style, v.Status))
	}
	return strings.TrimRight(b.String(), "\n")
}

// readKeys forwards learner commands. In raw TTY mode every byte is a key;
// otherwise each input line is a command ("h", "hint", "n", "quit", ...).
// On EOF in line mode it stops reading but the watcher keeps running.
func readKeys(r io.Reader, raw bool, keys chan<- key) {
	if raw {
		buf := make([]byte, 1)
		for {
			n, err := r.Read(buf)
			if err != nil {
				return
			}
			if n == 1 {
				keys <- key(toLower(buf[0]))
			}
		}
	}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.ToLower(strings.TrimSpace(sc.Text()))
		if line == "" {
			continue
		}
		keys <- key(line[0])
	}
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 'a' - 'A'
	}
	return b
}

// crlfWriter translates "\n" into "\r\n" for terminals in raw mode.
type crlfWriter struct{ w io.Writer }

func (c crlfWriter) Write(p []byte) (int, error) {
	if _, err := c.w.Write(bytes.ReplaceAll(p, []byte("\n"), []byte("\r\n"))); err != nil {
		return 0, err
	}
	return len(p), nil
}
