package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/maskedsyntax/golings/internal/app"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/exitcode"
	"github.com/maskedsyntax/golings/internal/progress"
	"github.com/maskedsyntax/golings/internal/render"
	"github.com/maskedsyntax/golings/internal/runner"
	"github.com/maskedsyntax/golings/internal/toolchain"
	"github.com/maskedsyntax/golings/internal/version"
	"github.com/maskedsyntax/golings/internal/workspace"
)

// exerciseJSON is the stable JSON shape of an exercise.
type exerciseJSON struct {
	ID               string   `json:"id"`
	Ordinal          int      `json:"ordinal"`
	Module           string   `json:"module"`
	Topic            string   `json:"topic"`
	Title            string   `json:"title"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	EstimatedMinutes int      `json:"estimated_minutes"`
	Files            []string `json:"files"`
	URL              string   `json:"url"`
	Status           string   `json:"status,omitempty"`
	Attempts         int      `json:"attempts"`
}

func toJSON(a *app.App, v app.View) exerciseJSON {
	ex := v.Exercise
	files := make([]string, len(ex.StarterFiles))
	for i, f := range ex.StarterFiles {
		files[i] = filepath.ToSlash(filepath.Join(a.Workspace.RelExerciseDir(ex), f))
	}
	return exerciseJSON{
		ID: ex.ID, Ordinal: ex.Ordinal, Module: ex.Module, Topic: ex.Topic, Title: ex.Title,
		Type: string(ex.Type), Difficulty: ex.Difficulty, EstimatedMinutes: ex.EstimatedMinutes,
		Files: files, URL: ex.URL, Status: string(v.Status), Attempts: v.State.Attempts,
	}
}

func (e *env) header(a *app.App, ex *catalog.Exercise) {
	topic := ex.Topic
	if t := a.Catalog.Topic(ex.Topic); t != nil {
		topic = t.Title
	}
	render.ExerciseHeader(e.stdout, e.style, ex, len(a.Catalog.Exercises()), topic, a.Workspace.RelExerciseDir(ex))
}

// resultCode maps an evaluation status to an exit code.
func resultCode(s runner.Status) int {
	switch s {
	case runner.StatusPass:
		return exitcode.OK
	case runner.StatusFail:
		return exitcode.VerifyFailed
	default:
		return exitcode.Infra
	}
}

func verifyCurrent(e *env) error {
	a, err := e.app()
	if err != nil {
		return err
	}
	st, err := a.Store.Load()
	if err != nil {
		return err
	}
	ex := a.Current(st)
	if ex == nil {
		if e.g.json {
			return render.JSON(e.stdout, map[string]any{"all_complete": true})
		}
		fmt.Fprintln(e.stdout, e.style.Green("🎉 Every exercise is complete. Well done!"))
		return nil
	}
	if st.CurrentExerciseID != ex.ID {
		if err := a.SelectCurrent(ex.ID); err != nil {
			return err
		}
	}
	return evaluateAndShow(e, a, ex, false)
}

// ---- init ------------------------------------------------------------------

func cmdInit() *command {
	var allowNonEmpty bool
	return &command{
		name:    "init",
		args:    "[path]",
		summary: "create a new workspace (default: ./" + workspace.DefaultDir + ")",
		setup: func(fs *flag.FlagSet) {
			fs.BoolVar(&allowNonEmpty, "force-into-nonempty", false, "allow a non-empty directory if no golings files would be overwritten")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) > 1 {
				return exitcode.New(exitcode.Usage, "init takes at most one path")
			}
			dest := workspace.DefaultDir
			if len(args) == 1 {
				dest = args[0]
			}
			cat, err := e.catalog()
			if err != nil {
				return err
			}
			ws, err := workspace.Init(dest, cat, workspace.InitOptions{AllowNonEmpty: allowNonEmpty, CLIVersion: version.String()})
			if errors.Is(err, workspace.ErrNotEmpty) {
				return exitcode.Wrap(exitcode.Workspace, fmt.Errorf("%w; choose another path or pass --force-into-nonempty", err))
			}
			if err != nil {
				return err
			}
			if e.g.json {
				return render.JSON(e.stdout, map[string]any{"workspace": ws.Root, "exercises": len(cat.Exercises()), "curriculum_version": cat.Version()})
			}
			fmt.Fprintf(e.stdout, "%s Created a workspace with %d exercises in %s\n\n", e.style.Green("✓"), len(cat.Exercises()), ws.Root)
			fmt.Fprintf(e.stdout, "Next steps:\n  cd %s\n  golings\n", quoteIfNeeded(dest))
			return nil
		},
	}
}

func quoteIfNeeded(s string) string {
	if strings.ContainsAny(s, " \t'\"") {
		return fmt.Sprintf("%q", s)
	}
	return s
}

// ---- list ------------------------------------------------------------------

func cmdList() *command {
	var module, topic, difficulty, status string
	return &command{
		name:    "list",
		summary: "list exercises and their status",
		setup: func(fs *flag.FlagSet) {
			fs.StringVar(&module, "module", "", "only this module (id or number, e.g. 01)")
			fs.StringVar(&topic, "topic", "", "only this topic (e.g. slices)")
			fs.StringVar(&difficulty, "difficulty", "", "beginner, intermediate or advanced")
			fs.StringVar(&status, "status", "", "not_started, in_progress, completed, modified, skipped or missing")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) > 0 {
				return exitcode.New(exitcode.Usage, "list takes no arguments")
			}
			a, err := e.app()
			if err != nil {
				return err
			}
			st, err := a.Store.Load()
			if err != nil {
				return err
			}
			var rows []app.View
			for _, v := range a.Views(st) {
				ex := v.Exercise
				switch {
				case module != "" && ex.Module != module && !strings.HasPrefix(ex.Module, module+"_"):
				case topic != "" && ex.Topic != topic:
				case difficulty != "" && ex.Difficulty != difficulty:
				case status != "" && string(v.Status) != status:
				default:
					rows = append(rows, v)
				}
			}
			if e.g.json {
				out := make([]exerciseJSON, 0, len(rows))
				for _, v := range rows {
					out = append(out, toJSON(a, v))
				}
				return render.JSON(e.stdout, out)
			}
			if len(rows) == 0 {
				fmt.Fprintln(e.stdout, "No exercises match.")
				return nil
			}
			current := a.Current(st)
			lastModule := ""
			for _, v := range rows {
				ex := v.Exercise
				if ex.Module != lastModule {
					if lastModule != "" {
						fmt.Fprintln(e.stdout)
					}
					fmt.Fprintln(e.stdout, e.style.Bold(fmt.Sprintf("%s  %s", ex.Module[:2], a.Catalog.ModuleTitle(ex.Module))))
					lastModule = ex.Module
				}
				marker := " "
				if current != nil && current.ID == ex.ID {
					marker = e.style.Cyan(">")
				}
				fmt.Fprintf(e.stdout, "%s %3d  %-*s %s  %s\n", marker, ex.Ordinal, a.Catalog.IDWidth(), ex.ID, statusCell(e.style, v.Status), ex.Title)
			}
			return nil
		},
	}
}

func statusCell(s render.Style, d app.DisplayStatus) string {
	label := fmt.Sprintf("%-11s", strings.ReplaceAll(string(d), "_", " "))
	switch d {
	case app.DisplayCompleted:
		return s.Green(label)
	case app.DisplayModified, app.DisplayInProgress, app.DisplayRecheck:
		return s.Yellow(label)
	case app.DisplayMissing:
		return s.Red(label)
	default:
		return s.Dim(label)
	}
}

// ---- run / verify ----------------------------------------------------------

func cmdRun() *command {
	return &command{
		name:    "run",
		args:    "<id>",
		summary: "run one exercise's checks once, showing program output",
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) != 1 {
				return exitcode.New(exitcode.Usage, "run needs exactly one exercise id (see `golings list`)")
			}
			a, err := e.app()
			if err != nil {
				return err
			}
			ex, err := a.Catalog.Find(args[0])
			if err != nil {
				return err
			}
			return evaluateAndShow(e, a, ex, true)
		},
	}
}

func cmdVerify() *command {
	var all bool
	return &command{
		name:    "verify",
		args:    "[id]",
		summary: "check one exercise (default: the current one), or every exercise with --all",
		setup: func(fs *flag.FlagSet) {
			fs.BoolVar(&all, "all", false, "verify every exercise in the workspace")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			switch {
			case all && len(args) > 0:
				return exitcode.New(exitcode.Usage, "use either an exercise id or --all, not both")
			case len(args) > 1:
				return exitcode.New(exitcode.Usage, "verify takes at most one exercise id")
			case all:
				return verifyAll(e)
			case len(args) == 1:
				a, err := e.app()
				if err != nil {
					return err
				}
				ex, err := a.Catalog.Find(args[0])
				if err != nil {
					return err
				}
				return evaluateAndShow(e, a, ex, false)
			default:
				return verifyCurrent(e)
			}
		},
	}
}

func evaluateAndShow(e *env, a *app.App, ex *catalog.Exercise, verbose bool) error {
	tc, err := e.toolchain()
	if err != nil {
		return err
	}
	if !e.g.json {
		e.header(a, ex)
		fmt.Fprintf(e.stdout, "\n%s\n\n", e.style.Dim("Checking…"))
	}
	ev, err := a.Evaluate(e.ctx, tc, ex, e.g.timeout)
	if err != nil {
		return err
	}
	if e.g.json {
		if err := render.JSON(e.stdout, map[string]any{"exercise": ex.ID, "result": ev.Result}); err != nil {
			return err
		}
		return exitFor(ev.Result.Status)
	}
	render.Result(e.stdout, e.style, ev.Result, verbose || e.g.verbose)
	if ev.Result.Passed() {
		if next := nextAfter(a, ex); next != nil {
			fmt.Fprintf(e.stdout, "Exercise complete! Run `golings next` to move on to %s.\n", next.ID)
		} else {
			fmt.Fprintln(e.stdout, "Exercise complete!")
		}
	} else if ev.Result.Status == runner.StatusFail {
		fmt.Fprintf(e.stdout, "%s\n", e.style.Dim("Stuck? Read "+filepath.Join(a.Workspace.RelExerciseDir(ex), "README.md")+"."))
	}
	return exitFor(ev.Result.Status)
}

func exitFor(s runner.Status) error {
	if code := resultCode(s); code != exitcode.OK {
		return silentExit(code)
	}
	return nil
}

func nextAfter(a *app.App, ex *catalog.Exercise) *catalog.Exercise {
	st, err := a.Store.Load()
	if err != nil {
		return nil
	}
	views := a.Views(st)
	for _, v := range views[ex.Ordinal:] {
		if !v.Status.Done() && v.Status != app.DisplaySkipped {
			return v.Exercise
		}
	}
	return a.NextUnfinished(st)
}

func verifyAll(e *env) error {
	a, err := e.app()
	if err != nil {
		return err
	}
	tc, err := e.toolchain()
	if err != nil {
		return err
	}
	// Exercises are evaluated in parallel (many failing concurrency
	// exercises only fail after a guard timeout), but reported in order.
	exs := a.Catalog.Exercises()
	type outcome struct {
		res *runner.ExerciseResult
		err error
	}
	outcomes := make([]outcome, len(exs))
	done := make([]chan struct{}, len(exs))
	for i := range done {
		done[i] = make(chan struct{})
	}
	ctx, cancel := context.WithCancel(e.ctx)
	defer cancel()
	jobs := make(chan int)
	go func() {
		defer close(jobs)
		for i := range exs {
			select {
			case jobs <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	for range max(1, runtime.NumCPU()/2) {
		go func() {
			for i := range jobs {
				ex := exs[i]
				ev, err := a.Evaluate(ctx, tc, ex, e.g.timeout)
				var missing *workspace.MissingFilesError
				switch {
				case errors.As(err, &missing):
					outcomes[i].res = &runner.ExerciseResult{
						ExerciseID: ex.ID, Status: runner.StatusFail,
						Checks: []runner.CheckResult{{Name: "files", Status: runner.StatusFail, Diagnostic: missing.Error()}},
					}
				case err != nil:
					outcomes[i].err = err
				default:
					outcomes[i].res = ev.Result
				}
				close(done[i])
			}
		}()
	}

	worst := runner.StatusPass
	var results []*runner.ExerciseResult
	passed := 0
	for i, ex := range exs {
		select {
		case <-done[i]:
		case <-ctx.Done():
			return ctx.Err()
		}
		if err := outcomes[i].err; err != nil {
			return err // the deferred cancel stops the remaining work
		}
		res := outcomes[i].res
		results = append(results, res)
		worst = runner.Worst(worst, res.Status)
		if res.Passed() {
			passed++
		}
		if !e.g.json {
			fmt.Fprintf(e.stdout, "%s %3d  %-*s %s\n", statusMark(e.style, res.Status), ex.Ordinal, a.Catalog.IDWidth(), ex.ID, problemSummary(res))
		}
	}
	if e.g.json {
		if err := render.JSON(e.stdout, map[string]any{"passed": passed, "total": len(results), "results": results}); err != nil {
			return err
		}
	} else {
		fmt.Fprintf(e.stdout, "\n%d of %d exercises pass.\n", passed, len(results))
	}
	return exitFor(worst)
}

func statusMark(s render.Style, st runner.Status) string {
	switch st {
	case runner.StatusPass:
		return s.Green("✓")
	case runner.StatusUnsupported:
		return s.Yellow("-")
	default:
		return s.Red("✗")
	}
}

func problemSummary(res *runner.ExerciseResult) string {
	c := res.FirstProblem()
	if c == nil {
		return ""
	}
	if c.Failure != "" {
		return string(c.Failure)
	}
	return string(c.Status)
}

// ---- next ------------------------------------------------------------------

func cmdNext() *command {
	return &command{
		name:    "next",
		summary: "select and show the next unfinished exercise",
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) > 0 {
				return exitcode.New(exitcode.Usage, "next takes no arguments")
			}
			a, err := e.app()
			if err != nil {
				return err
			}
			st, err := a.Store.Load()
			if err != nil {
				return err
			}
			next := pickNext(a, st)
			if next == nil {
				if e.g.json {
					return render.JSON(e.stdout, map[string]any{"all_complete": true})
				}
				fmt.Fprintln(e.stdout, e.style.Green("🎉 Every exercise is complete. Well done!"))
				return nil
			}
			if err := a.SelectCurrent(next.ID); err != nil {
				return err
			}
			if e.g.json {
				return render.JSON(e.stdout, toJSON(a, a.ViewOf(next, st)))
			}
			e.header(a, next)
			var missing []string
			for _, p := range next.Prerequisites {
				if pre, ok := a.Catalog.Get(p); ok && !a.ViewOf(pre, st).Status.Done() {
					missing = append(missing, p)
				}
			}
			if len(missing) > 0 {
				fmt.Fprintf(e.stdout, "%s you haven't finished %s yet, which this exercise builds on.\n",
					e.style.Yellow("note:"), strings.Join(missing, ", "))
			}
			fmt.Fprintf(e.stdout, "\nOpen the file in your editor, then run `golings` to check it.\n")
			return nil
		},
	}
}

// pickNext returns the first unfinished exercise after the current one,
// wrapping around to the first unfinished exercise overall. Even when the
// current exercise is unfinished, "next" means "move on".
func pickNext(a *app.App, st *progress.State) *catalog.Exercise {
	views := a.Views(st)
	start := 0
	if cur, ok := a.Catalog.Get(st.CurrentExerciseID); ok {
		start = cur.Ordinal // index of the exercise after the current one
	}
	for i := range views {
		v := views[(start+i)%len(views)]
		if !v.Status.Done() && v.Status != app.DisplaySkipped {
			return v.Exercise
		}
	}
	return nil
}

// ---- progress --------------------------------------------------------------

func cmdProgress() *command {
	return &command{
		name:    "progress",
		summary: "show completion statistics",
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) > 0 {
				return exitcode.New(exitcode.Usage, "progress takes no arguments")
			}
			a, err := e.app()
			if err != nil {
				return err
			}
			st, err := a.Store.Load()
			if err != nil {
				return err
			}
			sum := a.Summarize(st, runtime.GOOS)
			if e.g.json {
				return render.JSON(e.stdout, sum)
			}
			fmt.Fprintf(e.stdout, "%s %d/%d exercises complete (%.0f%%)\n", e.style.Bold("Progress:"), sum.Completed, sum.Total, sum.Percent)
			fmt.Fprintf(e.stdout, "%s\n\n", progressBar(sum.Completed, sum.Total, 40))
			for _, m := range sum.Modules {
				fmt.Fprintf(e.stdout, "  %-48s %d/%d\n", m.Title, m.Completed, m.Total)
			}
			if sum.Modified > 0 {
				fmt.Fprintf(e.stdout, "\n%s %d completed exercise(s) were edited after passing; run `golings verify --all` to re-check them.\n", e.style.Yellow("note:"), sum.Modified)
			}
			if sum.Recheck > 0 {
				fmt.Fprintf(e.stdout, "\n%s the checks of %d completed exercise(s) changed in a curriculum update; verify them again (`golings next` takes you there).\n", e.style.Yellow("note:"), sum.Recheck)
			}
			if sum.Unsupported > 0 {
				fmt.Fprintf(e.stdout, "\n%d exercise(s) are not available on %s and are not counted.\n", sum.Unsupported, runtime.GOOS)
			}
			return nil
		},
	}
}

func progressBar(done, total, width int) string {
	if total == 0 {
		return "[" + strings.Repeat(" ", width) + "]"
	}
	n := done * width / total
	return "[" + strings.Repeat("#", n) + strings.Repeat("-", width-n) + "]"
}

// ---- doctor ----------------------------------------------------------------

type doctorCheck struct {
	Name   string `json:"name"`
	OK     bool   `json:"ok"`
	Detail string `json:"detail"`
}

func cmdDoctor() *command {
	var recoverState bool
	return &command{
		name:    "doctor",
		summary: "diagnose the Go toolchain, the workspace and the platform",
		setup: func(fs *flag.FlagSet) {
			fs.BoolVar(&recoverState, "recover-state", false, "move a corrupt progress file aside and start fresh")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) > 0 {
				return exitcode.New(exitcode.Usage, "doctor takes no arguments")
			}
			var checks []doctorCheck
			code := exitcode.OK
			add := func(name string, ok bool, detail string, failCode int) {
				checks = append(checks, doctorCheck{name, ok, detail})
				if !ok && code == exitcode.OK {
					code = failCode
				}
			}

			tc, err := toolchain.Detect(e.ctx)
			switch {
			case err != nil:
				add("go toolchain", false, err.Error(), exitcode.Toolchain)
			default:
				add("go toolchain", true, fmt.Sprintf("%s at %s (%s/%s)", tc.Version, tc.GoBin, tc.GOOS, tc.GOARCH), 0)
				race := "available"
				if !tc.RaceSupported {
					race = "not available (concurrency exercises with race checks cannot be fully verified)"
				}
				checks = append(checks, doctorCheck{"race detector", tc.RaceSupported, race})
			}

			if dir, err := os.MkdirTemp("", "golings-doctor-*"); err != nil {
				add("temp directory", false, err.Error(), exitcode.Infra)
			} else {
				os.RemoveAll(dir)
				add("temp directory", true, os.TempDir(), 0)
			}

			if _, err := e.catalog(); err != nil {
				add("curriculum", false, err.Error(), exitcode.Infra)
			} else {
				add("curriculum", true, fmt.Sprintf("version %s, %d exercises", e.cat.Version(), len(e.cat.Exercises())), 0)
			}

			ws, err := e.workspace()
			switch {
			case errors.Is(err, workspace.ErrNotWorkspace) && e.g.workspace == "":
				checks = append(checks, doctorCheck{"workspace", true, "not inside a workspace (run `golings init` to create one)"})
			case err != nil:
				add("workspace", false, err.Error(), exitcode.Workspace)
			default:
				add("workspace", true, ws.Root, 0)
				store := progress.NewStore(ws.MetaPath(), ws.Meta.CurriculumVersion)
				if recoverState {
					moved, err := store.Recover()
					switch {
					case err != nil:
						add("progress", false, err.Error(), exitcode.Workspace)
					case moved != "":
						add("progress", true, "moved the old progress file to "+moved+"; starting fresh", 0)
					default:
						add("progress", true, "nothing to recover", 0)
					}
				} else if _, err := store.Load(); err != nil {
					add("progress", false, err.Error(), exitcode.Workspace)
				} else {
					add("progress", true, store.Path(), 0)
				}
				probe := filepath.Join(ws.MetaPath(), ".write-test")
				if err := os.WriteFile(probe, nil, 0o644); err != nil {
					add("workspace writable", false, err.Error(), exitcode.Workspace)
				} else {
					os.Remove(probe)
					add("workspace writable", true, "ok", 0)
				}
			}

			if e.g.json {
				if err := render.JSON(e.stdout, map[string]any{"ok": code == exitcode.OK, "checks": checks}); err != nil {
					return err
				}
			} else {
				for _, c := range checks {
					mark := e.style.Green("✓")
					if !c.OK {
						mark = e.style.Red("✗")
					}
					fmt.Fprintf(e.stdout, "%s %-20s %s\n", mark, c.Name, c.Detail)
				}
			}
			if code != exitcode.OK {
				return silentExit(code)
			}
			return nil
		},
	}
}

// ---- version / help --------------------------------------------------------

func cmdVersion() *command {
	return &command{
		name:    "version",
		summary: "print golings and curriculum versions",
		run: func(e *env, _ *flag.FlagSet, _ []string) error {
			cat, err := e.catalog()
			if err != nil {
				return err
			}
			if e.g.json {
				return render.JSON(e.stdout, map[string]any{
					"golings": version.String(), "curriculum": cat.Version(),
					"exercises": len(cat.Exercises()), "built_with": runtime.Version(),
					"workspace_format": workspace.FormatVersion,
				})
			}
			fmt.Fprintf(e.stdout, "golings %s (curriculum %s, %d exercises, built with %s)\n",
				version.String(), cat.Version(), len(cat.Exercises()), runtime.Version())
			return nil
		},
	}
}

func cmdHelp() *command {
	return &command{
		name:    "help",
		args:    "[command]",
		summary: "show help",
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) == 0 {
				printUsage(e.stdout)
				return nil
			}
			c := lookup(args[0])
			if c == nil {
				return exitcode.New(exitcode.Usage, "unknown command %q", args[0])
			}
			fs := flag.NewFlagSet(c.name, flag.ContinueOnError)
			if c.setup != nil {
				c.setup(fs)
			}
			printCommandUsage(e.stdout, c, fs)
			return nil
		},
	}
}
