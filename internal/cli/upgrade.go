package cli

import (
	"errors"
	"flag"
	"fmt"
	"slices"

	"github.com/maskedsyntax/golings/internal/app"
	"github.com/maskedsyntax/golings/internal/exitcode"
	"github.com/maskedsyntax/golings/internal/render"
	"github.com/maskedsyntax/golings/internal/workspace"
)

func cmdUpgrade() *command {
	var dryRun bool
	return &command{
		name:    "upgrade",
		summary: "bring the workspace up to this golings' curriculum (your files are never rewritten)",
		setup: func(fs *flag.FlagSet) {
			fs.BoolVar(&dryRun, "dry-run", false, "show what would change without changing anything")
		},
		run: func(e *env, _ *flag.FlagSet, args []string) error {
			if len(args) > 0 {
				return exitcode.New(exitcode.Usage, "upgrade takes no arguments")
			}
			cat, err := e.catalog()
			if err != nil {
				return err
			}
			ws, err := e.workspace()
			if err != nil {
				return err
			}
			res, err := app.New(cat, ws).Upgrade(dryRun)
			if errors.Is(err, workspace.ErrWorkspaceNewer) {
				return exitcode.Wrap(exitcode.Workspace, err)
			}
			if err != nil {
				return err
			}
			if e.g.json {
				if err := render.JSON(e.stdout, res); err != nil {
					return err
				}
			} else {
				printUpgrade(e, res)
			}
			if len(res.Conflicts) > 0 {
				return exitcode.New(exitcode.Workspace, "upgrade blocked; resolve the conflicts above and run it again")
			}
			return nil
		},
	}
}

func printUpgrade(e *env, r *app.UpgradeResult) {
	w, s := e.stdout, e.style
	if r.FromVersion == r.ToVersion && r.Empty() && len(r.ProgressRenamed) == 0 && len(r.Conflicts) == 0 {
		fmt.Fprintf(w, "%s Workspace is up to date (curriculum %s).\n", s.Green("✓"), r.ToVersion)
		return
	}
	verb := "Upgraded"
	if r.DryRun {
		verb = "Would upgrade"
	}
	if len(r.Conflicts) > 0 {
		verb = "Cannot upgrade"
	}
	fmt.Fprintf(w, "%s workspace from curriculum %s to %s.\n", s.Bold(verb), r.FromVersion, r.ToVersion)
	list := func(title string, items []string) {
		if len(items) == 0 {
			return
		}
		fmt.Fprintf(w, "\n%s (%d):\n", title, len(items))
		for _, it := range items {
			fmt.Fprintf(w, "  %s\n", it)
		}
	}
	list("New exercises", r.Added)
	var moves []string
	for _, m := range r.Moved {
		moves = append(moves, fmt.Sprintf("%s -> %s", m.From, m.To))
	}
	list("Moved, with your files", moves)
	var renamed []string
	for old, cur := range r.ProgressRenamed {
		renamed = append(renamed, old+" -> "+cur)
	}
	list("Progress carried over", sortedStrings(renamed))
	if n := len(r.Refreshed); n > 0 {
		fmt.Fprintf(w, "\nUpdated README or data files of %d exercise(s).\n", n)
	}
	list("No longer in the curriculum (left untouched)", r.Retired)
	list(s.Red("Conflicts"), r.Conflicts)
	if r.StateBackup != "" {
		fmt.Fprintf(w, "\nProgress was backed up to %s\n", r.StateBackup)
	}
	if r.DryRun && len(r.Conflicts) == 0 {
		fmt.Fprintf(w, "\nRun `golings upgrade` to apply.\n")
	}
}

func sortedStrings(s []string) []string {
	slices.Sort(s)
	return s
}
