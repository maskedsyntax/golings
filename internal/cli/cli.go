// Package cli implements the golings command line: argument parsing,
// terminal output and exit codes. Evaluation logic lives in package app.
package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/app"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/exitcode"
	"github.com/maskedsyntax/golings/internal/progress"
	"github.com/maskedsyntax/golings/internal/render"
	"github.com/maskedsyntax/golings/internal/toolchain"
	"github.com/maskedsyntax/golings/internal/workspace"
)

// globals are flags accepted by every command.
type globals struct {
	workspace string
	noColor   bool
	json      bool
	verbose   bool
	timeout   time.Duration
}

func (g *globals) register(fs *flag.FlagSet) {
	fs.StringVar(&g.workspace, "workspace", "", "path to the workspace (default: search upward from the current directory)")
	fs.BoolVar(&g.noColor, "no-color", false, "disable colored output (NO_COLOR is also honored)")
	fs.BoolVar(&g.json, "json", false, "print machine-readable JSON where supported")
	fs.BoolVar(&g.verbose, "verbose", false, "show more detail")
	fs.DurationVar(&g.timeout, "timeout", 0, "override the per-exercise time limit (e.g. 30s)")
}

// env carries per-invocation state.
type env struct {
	ctx    context.Context
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	g      globals
	style  render.Style

	cat *catalog.Catalog
}

type command struct {
	name    string
	args    string
	summary string
	setup   func(fs *flag.FlagSet) // registers command-specific flags
	run     func(e *env, fs *flag.FlagSet, args []string) error
}

var commands []*command

func init() {
	commands = []*command{
		cmdWatch(), cmdInit(), cmdList(), cmdRun(), cmdVerify(), cmdNext(), cmdHint(),
		cmdReset(), cmdSolution(), cmdProgress(), cmdUpgrade(), cmdDoctor(), cmdVersion(), cmdHelp(),
	}
}

func lookup(name string) *command {
	for _, c := range commands {
		if c.name == name {
			return c
		}
	}
	return nil
}

// Main runs the CLI and returns the process exit code.
func Main(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return run(ctx, args, stdin, stdout, stderr)
}

func run(ctx context.Context, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	e := &env{ctx: ctx, stdin: stdin, stdout: stdout, stderr: stderr}

	// Global flags may precede the command name.
	pre := flag.NewFlagSet("golings", flag.ContinueOnError)
	pre.SetOutput(io.Discard)
	e.g.register(pre)
	if err := pre.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			printUsage(stdout)
			return exitcode.OK
		}
		fmt.Fprintf(stderr, "golings: %v\n\n", err)
		printUsage(stderr)
		return exitcode.Usage
	}
	rest := pre.Args()

	name := "watch"
	if len(rest) > 0 {
		name, rest = rest[0], rest[1:]
	}
	cmd := lookup(name)
	if cmd == nil {
		fmt.Fprintf(stderr, "golings: unknown command %q\n\n", name)
		printUsage(stderr)
		return exitcode.Usage
	}

	// Remember global flags given before the command name: registering them
	// again on the command's flag set resets the shared fields to defaults.
	preset := map[string]string{}
	pre.Visit(func(f *flag.Flag) { preset[f.Name] = f.Value.String() })

	fs := flag.NewFlagSet("golings "+cmd.name, flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	e.g.register(fs)
	for name, value := range preset {
		fs.Set(name, value)
	}
	if cmd.setup != nil {
		cmd.setup(fs)
	}
	positional, err := parseInterspersed(fs, rest)
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			printCommandUsage(stdout, cmd, fs)
			return exitcode.OK
		}
		fmt.Fprintf(stderr, "golings %s: %v\n", cmd.name, err)
		printCommandUsage(stderr, cmd, fs)
		return exitcode.Usage
	}
	e.style = render.Style{Color: !e.g.json && render.ColorEnabled(stdout, e.g.noColor)}

	err = cmd.run(e, fs, positional)
	if err == nil {
		return exitcode.OK
	}
	var silent silentExit
	if errors.As(err, &silent) {
		return int(silent)
	}
	if app.IsCanceled(err) && ctx.Err() != nil {
		fmt.Fprintln(stderr, "golings: interrupted")
		return exitcode.Interrupted
	}
	fmt.Fprintf(stderr, "golings: %v\n", err)
	return classify(err)
}

// silentExit ends a command with a code after it already printed its output.
type silentExit int

func (s silentExit) Error() string { return fmt.Sprintf("exit status %d", int(s)) }

// classify maps well-known errors to exit codes.
func classify(err error) int {
	var (
		ee       *exitcode.Error
		notFound *catalog.NotFoundError
		ambig    *catalog.AmbiguousError
		corrupt  *progress.CorruptError
		missing  *workspace.MissingFilesError
		tooOld   *toolchain.TooOldError
	)
	switch {
	case errors.As(err, &ee):
		return ee.Code
	case errors.As(err, &notFound), errors.As(err, &ambig):
		return exitcode.Usage
	case errors.Is(err, workspace.ErrNotWorkspace), errors.As(err, &corrupt), errors.As(err, &missing),
		errors.Is(err, workspace.ErrNotEmpty), errors.Is(err, workspace.ErrAlreadyWorkspace):
		return exitcode.Workspace
	case errors.Is(err, toolchain.ErrNotFound), errors.As(err, &tooOld):
		return exitcode.Toolchain
	default:
		return exitcode.Infra
	}
}

// parseInterspersed parses flags that may appear before, between, or after
// positional arguments. Everything after a literal "--" is positional.
func parseInterspersed(fs *flag.FlagSet, args []string) ([]string, error) {
	var tail []string
	for i, a := range args {
		if a == "--" {
			args, tail = args[:i], args[i+1:]
			break
		}
	}
	var positional []string
	for {
		if err := fs.Parse(args); err != nil {
			return nil, err
		}
		args = fs.Args()
		if len(args) == 0 {
			return append(positional, tail...), nil
		}
		positional = append(positional, args[0])
		args = args[1:]
	}
}

func (e *env) catalog() (*catalog.Catalog, error) {
	if e.cat != nil {
		return e.cat, nil
	}
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		return nil, exitcode.Wrap(exitcode.Infra, fmt.Errorf("embedded curriculum is invalid (this is a golings bug): %w", err))
	}
	e.cat = cat
	return cat, nil
}

func (e *env) workspace() (*workspace.Workspace, error) {
	if e.g.workspace != "" {
		return workspace.Open(e.g.workspace)
	}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return workspace.Discover(wd)
}

// app opens the catalog and workspace and checks that they match.
func (e *env) app() (*app.App, error) {
	cat, err := e.catalog()
	if err != nil {
		return nil, err
	}
	ws, err := e.workspace()
	if err != nil {
		return nil, err
	}
	if ws.Meta.CurriculumVersion != cat.Version() && !e.g.json {
		fmt.Fprintf(e.stderr, "%s workspace uses curriculum %s; this golings ships %s. Run `golings upgrade` to update it.\n",
			e.style.Yellow("note:"), ws.Meta.CurriculumVersion, cat.Version())
	}
	return app.New(cat, ws), nil
}

func (e *env) toolchain() (*toolchain.Info, error) {
	tc, err := toolchain.Detect(e.ctx)
	if err != nil {
		return nil, exitcode.Wrap(exitcode.Toolchain, err)
	}
	return tc, nil
}

func printUsage(w io.Writer) {
	fmt.Fprint(w, `golings: learn Go by fixing Go.

Usage:
  golings [command] [flags] [arguments]

Commands:
`)
	for _, c := range commands {
		fmt.Fprintf(w, "  %-28s %s\n", strings.TrimSpace(c.name+" "+c.args), c.summary)
	}
	fmt.Fprint(w, `
Global flags:
  --workspace <path>   use this workspace instead of searching upward
  --no-color           disable colors (NO_COLOR is honored too)
  --json               machine-readable output (list, verify, next, progress, upgrade, doctor, version)
  --verbose            show more detail
  --timeout <dur>      override the per-exercise time limit

Exit codes: 0 ok, 1 exercise not passing, 2 usage error, 3 workspace problem,
4 Go toolchain problem, 5 infrastructure error or time limit, 130 interrupted.

Run "golings help <command>" for details on one command.
`)
}

func printCommandUsage(w io.Writer, c *command, fs *flag.FlagSet) {
	fmt.Fprintf(w, "Usage: golings %s\n\n%s\n", strings.TrimSpace(c.name+" "+c.args), c.summary)
	var local []string
	fs.VisitAll(func(f *flag.Flag) {
		switch f.Name {
		case "workspace", "no-color", "json", "verbose", "timeout":
			return
		}
		local = append(local, fmt.Sprintf("  --%-20s %s", f.Name, f.Usage))
	})
	if len(local) > 0 {
		fmt.Fprintf(w, "\nFlags:\n%s\n", strings.Join(local, "\n"))
	}
	fmt.Fprintln(w, "\nGlobal flags: --workspace, --no-color, --json, --verbose, --timeout (see `golings help`).")
}
