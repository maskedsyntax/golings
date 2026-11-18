package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/maskedsyntax/golings/internal/app"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/exitcode"
	"github.com/maskedsyntax/golings/internal/render"
	"github.com/maskedsyntax/golings/internal/workspace"
)

// interactive reports whether we can prompt: both stdin and stdout must be
// terminals, and JSON mode never prompts.
func (e *env) interactive() bool {
	if e.g.json {
		return false
	}
	in, ok1 := e.stdin.(*os.File)
	out, ok2 := e.stdout.(*os.File)
	return ok1 && ok2 && render.IsTerminal(in) && render.IsTerminal(out)
}

// confirm asks a yes/no question; the default is no.
func (e *env) confirm(question string) bool {
	fmt.Fprintf(e.stdout, "%s [y/N] ", question)
	line, _ := bufio.NewReader(e.stdin).ReadString('\n')
	switch strings.ToLower(strings.TrimSpace(line)) {
	case "y", "yes":
		return true
	}
	return false
}

// resolve finds the exercise named by args[0], or the current exercise.
func (e *env) resolve(a *app.App, args []string, cmd string) (*catalog.Exercise, error) {
	if len(args) > 1 {
		return nil, exitcode.New(exitcode.Usage, "%s takes at most one exercise id", cmd)
	}
	if len(args) == 1 {
		return a.Catalog.Find(args[0])
	}
	st, err := a.Store.Load()
	if err != nil {
		return nil, err
	}
	ex := a.Current(st)
	if ex == nil {
		return nil, exitcode.New(exitcode.Usage, "every exercise is complete; name one explicitly")
	}
	return ex, nil
}

// ---- hint ------------------------------------------------------------------

func cmdHint() *command {
	var level int
	return &command{
		name:    "hint",
		args:    "[id]",
		summary: "reveal the next hint for an exercise (default: the current one)",
		setup: func(fs *flag.FlagSet) {
			fs.IntVar(&level, "level", 0, "show this hint level (1-3) instead of the next one")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			a, err := e.app()
			if err != nil {
				return err
			}
			ex, err := e.resolve(a, args, "hint")
			if err != nil {
				return err
			}
			h, err := a.RevealHint(ex, level)
			if err != nil {
				return exitcode.Wrap(exitcode.Usage, err)
			}
			if e.g.json {
				return render.JSON(e.stdout, map[string]any{"exercise": ex.ID, "hint": h})
			}
			printHint(e, ex, h)
			return nil
		},
	}
}

func printHint(e *env, ex *catalog.Exercise, h app.Hint) {
	fmt.Fprintf(e.stdout, "%s %s\n", e.style.Yellow(fmt.Sprintf("Hint %d/%d for %s:", h.Level, h.Total, ex.ID)), h.Text)
	if h.AllRevealed {
		fmt.Fprintf(e.stdout, "%s\n", e.style.Dim("That was the last hint. If you're still stuck, `golings solution "+ex.ID+" --show` reveals the answer."))
	}
}

// ---- reset -----------------------------------------------------------------

func cmdReset() *command {
	var force, noBackup bool
	return &command{
		name:    "reset",
		args:    "<id>",
		summary: "restore an exercise's original starter files (your version is backed up)",
		setup: func(fs *flag.FlagSet) {
			fs.BoolVar(&force, "force", false, "do not ask for confirmation")
			fs.BoolVar(&noBackup, "no-backup", false, "do not keep a backup of your current files")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) != 1 {
				return exitcode.New(exitcode.Usage, "reset needs exactly one exercise id")
			}
			a, err := e.app()
			if err != nil {
				return err
			}
			ex, err := a.Catalog.Find(args[0])
			if err != nil {
				return err
			}
			plan, err := a.Workspace.PlanReset(a.Catalog, ex)
			if err != nil {
				return err
			}
			rel := a.Workspace.RelExerciseDir(ex)
			if plan.Empty() {
				if e.g.json {
					return render.JSON(e.stdout, map[string]any{"exercise": ex.ID, "changed": false})
				}
				fmt.Fprintf(e.stdout, "%s already matches its starter files; nothing to reset.\n", ex.ID)
				return nil
			}
			if !force {
				if !e.interactive() {
					return exitcode.New(exitcode.Usage, "reset would replace files in %s; pass --force to confirm when not running interactively", rel)
				}
				fmt.Fprintf(e.stdout, "Resetting %s will:\n", ex.ID)
				for _, f := range plan.Changed {
					fmt.Fprintf(e.stdout, "  replace  %s\n", filepath.Join(rel, f))
				}
				for _, f := range plan.Missing {
					fmt.Fprintf(e.stdout, "  restore  %s\n", filepath.Join(rel, f))
				}
				if !noBackup && len(plan.Changed) > 0 {
					fmt.Fprintln(e.stdout, "Your current files will be backed up under .golings/backups/.")
				}
				if !e.confirm("Continue?") {
					fmt.Fprintln(e.stdout, "Reset canceled; nothing changed.")
					return nil
				}
			}
			backup, err := a.Workspace.Reset(a.Catalog, ex, !noBackup)
			if err != nil {
				return err
			}
			if e.g.json {
				return render.JSON(e.stdout, map[string]any{"exercise": ex.ID, "changed": true, "backup": backup})
			}
			fmt.Fprintf(e.stdout, "%s Restored the starter files for %s.\n", e.style.Green("✓"), ex.ID)
			if backup != "" {
				if relBackup, err := filepath.Rel(a.Workspace.Root, backup); err == nil {
					backup = relBackup
				}
				fmt.Fprintf(e.stdout, "Your previous version is in %s\n", backup)
			}
			return nil
		},
	}
}

// ---- solution --------------------------------------------------------------

func cmdSolution() *command {
	var show bool
	var out string
	return &command{
		name:    "solution",
		args:    "<id>",
		summary: "show the reference solution (asks first, to avoid accidental spoilers)",
		setup: func(fs *flag.FlagSet) {
			fs.BoolVar(&show, "show", false, "show the solution without asking")
			fs.StringVar(&out, "out", "", "write the solution files into this new directory instead of printing them")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) != 1 {
				return exitcode.New(exitcode.Usage, "solution needs exactly one exercise id")
			}
			a, err := e.app()
			if err != nil {
				return err
			}
			ex, err := a.Catalog.Find(args[0])
			if err != nil {
				return err
			}
			if !show {
				if !e.interactive() {
					return exitcode.New(exitcode.Usage, "pass --show to display the solution for %s", ex.ID)
				}
				if !e.confirm(fmt.Sprintf("Show the solution for %s? It gives the answer away.", ex.ID)) {
					return nil
				}
			}
			files, err := workspace.SolutionFiles(a.Catalog, ex)
			if err != nil {
				return err
			}
			if out != "" {
				if err := writeSolutionDir(out, files); err != nil {
					return err
				}
			}
			if err := a.MarkSolutionViewed(ex); err != nil {
				return err
			}
			names := make([]string, 0, len(files))
			for n := range files {
				names = append(names, n)
			}
			sort.Strings(names)
			if e.g.json {
				m := map[string]string{}
				for _, n := range names {
					m[n] = string(files[n])
				}
				return render.JSON(e.stdout, map[string]any{"exercise": ex.ID, "files": m, "written_to": out})
			}
			if out != "" {
				fmt.Fprintf(e.stdout, "%s Wrote the solution for %s to %s\n", e.style.Green("✓"), ex.ID, out)
				return nil
			}
			for _, n := range names {
				fmt.Fprintf(e.stdout, "%s\n%s\n", e.style.Bold("// ---- "+n+" ----"), strings.TrimRight(string(files[n]), "\n"))
			}
			fmt.Fprintf(e.stdout, "\n%s\n", e.style.Dim("Your files were not changed. Viewing a solution doesn't mark the exercise complete."))
			return nil
		},
	}
}

// writeSolutionDir writes files into dir, which must not exist yet or be
// empty, so learner code is never overwritten.
func writeSolutionDir(dir string, files map[string][]byte) error {
	entries, err := os.ReadDir(dir)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	case err != nil:
		return err
	case len(entries) > 0:
		return exitcode.New(exitcode.Usage, "%s is not empty; choose a new directory for --out", dir)
	}
	for name, data := range files {
		f, err := os.OpenFile(filepath.Join(dir, name), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if err != nil {
			return err
		}
		if _, err := f.Write(data); err != nil {
			f.Close()
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	}
	return nil
}
