// Package app implements golings use cases on top of the catalog,
// workspace, runner and progress packages. It contains no terminal I/O.
package app

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/progress"
	"github.com/maskedsyntax/golings/internal/runner"
	"github.com/maskedsyntax/golings/internal/toolchain"
	"github.com/maskedsyntax/golings/internal/workspace"
)

// App binds a catalog to a workspace.
type App struct {
	Catalog   *catalog.Catalog
	Workspace *workspace.Workspace
	Store     *progress.Store
	// GOOS is the platform exercises are checked on. Exercises that don't
	// support it are never chosen automatically as the next one to do.
	GOOS string
}

// New returns an App for an opened workspace.
func New(cat *catalog.Catalog, ws *workspace.Workspace) *App {
	return &App{Catalog: cat, Workspace: ws, Store: progress.NewStore(ws.MetaPath(), cat.Version()), GOOS: runtime.GOOS}
}

// pending reports whether v still needs doing on this platform.
func (a *App) pending(v View) bool {
	if v.Status.Done() || v.Status == DisplaySkipped {
		return false
	}
	return a.GOOS == "" || v.Exercise.SupportsOS(a.GOOS)
}

// DisplayStatus is the status shown to learners. It refines the persisted
// workflow status with what is on disk right now.
type DisplayStatus string

const (
	DisplayNotStarted DisplayStatus = "not_started"
	DisplayInProgress DisplayStatus = "in_progress"
	DisplayCompleted  DisplayStatus = "completed"
	DisplayModified   DisplayStatus = "modified" // completed, but edited since the passing run
	DisplaySkipped    DisplayStatus = "skipped"
	DisplayMissing    DisplayStatus = "missing" // learner files are absent
)

// Done reports whether the status counts toward completion.
func (d DisplayStatus) Done() bool { return d == DisplayCompleted || d == DisplayModified }

// View is an exercise with its progress and on-disk status.
type View struct {
	Exercise *catalog.Exercise      `json:"-"`
	State    progress.ExerciseState `json:"-"`
	Status   DisplayStatus          `json:"status"`
}

// ViewOf computes the display status of ex. A missing exercise is never
// counted as complete, whatever state.json says.
func (a *App) ViewOf(ex *catalog.Exercise, st *progress.State) View {
	es := st.Get(ex.ID)
	v := View{Exercise: ex, State: es}
	files, err := a.Workspace.ReadLearnerFiles(ex)
	switch {
	case err != nil:
		v.Status = DisplayMissing
	case es.Status == progress.Completed && es.SourceFingerprint != workspace.Fingerprint(files):
		v.Status = DisplayModified
	case es.Status == progress.Completed:
		v.Status = DisplayCompleted
	case es.Status == progress.Skipped:
		v.Status = DisplaySkipped
	case es.Status == progress.InProgress:
		v.Status = DisplayInProgress
	default:
		v.Status = DisplayNotStarted
	}
	return v
}

// Views returns a view of every exercise in order.
func (a *App) Views(st *progress.State) []View {
	out := make([]View, 0, len(a.Catalog.Exercises()))
	for _, ex := range a.Catalog.Exercises() {
		out = append(out, a.ViewOf(ex, st))
	}
	return out
}

// NextUnfinished returns the first exercise that is neither completed nor
// skipped, or nil when everything is done.
func (a *App) NextUnfinished(st *progress.State) *catalog.Exercise {
	for _, v := range a.Views(st) {
		if a.pending(v) {
			return v.Exercise
		}
	}
	return nil
}

// Current returns the selected exercise if it is still unfinished, otherwise
// the next unfinished one. It returns nil when everything is done.
func (a *App) Current(st *progress.State) *catalog.Exercise {
	if ex, ok := a.Catalog.Get(st.CurrentExerciseID); ok {
		if a.pending(a.ViewOf(ex, st)) {
			return ex
		}
	}
	return a.NextUnfinished(st)
}

// SelectCurrent persists id as the current exercise.
func (a *App) SelectCurrent(id string) error {
	_, err := a.Store.Update(func(s *progress.State) error {
		s.CurrentExerciseID = id
		return nil
	})
	return err
}

// Evaluation is the outcome of Evaluate.
type Evaluation struct {
	Exercise *catalog.Exercise
	Result   *runner.ExerciseResult
	// NewlyCompleted is true when this run completed an exercise that was
	// not completed (or was modified) before.
	NewlyCompleted bool
}

// Evaluate reads the learner's files, runs the checks, and records the
// result. Progress is updated only after evaluation finishes; a canceled
// run records nothing.
func (a *App) Evaluate(ctx context.Context, tc *toolchain.Info, ex *catalog.Exercise, timeout time.Duration) (*Evaluation, error) {
	files, err := a.Workspace.ReadLearnerFiles(ex)
	if err != nil {
		return nil, err
	}
	assets, err := a.Catalog.Assets(ex)
	if err != nil {
		return nil, err
	}
	res := runner.Run(ctx, runner.Request{
		Exercise:   ex,
		Files:      files,
		Assets:     assets,
		Toolchain:  tc,
		Timeout:    timeout,
		DisplayDir: a.Workspace.RelExerciseDir(ex),
	})
	ev := &Evaluation{Exercise: ex, Result: res}
	if ctx.Err() != nil {
		return ev, ctx.Err()
	}
	// Runner errors and unsupported platforms say nothing about the
	// learner's work, so they are not counted as attempts.
	if res.Status == runner.StatusRunnerError || res.Status == runner.StatusUnsupported {
		return ev, nil
	}
	fp := workspace.Fingerprint(files)
	_, err = a.Store.Update(func(s *progress.State) error {
		before := s.Get(ex.ID)
		ev.NewlyCompleted = res.Passed() && (before.Status != progress.Completed || before.SourceFingerprint != fp)
		s.RecordResult(ex.ID, res.Passed(), string(res.Status), fp, time.Now())
		return nil
	})
	if err != nil {
		return ev, fmt.Errorf("recording progress: %w", err)
	}
	return ev, nil
}

// Summary is aggregate progress.
type Summary struct {
	Total       int             `json:"total"`
	Completed   int             `json:"completed"`
	Modified    int             `json:"modified_since_verified"`
	Skipped     int             `json:"skipped"`
	Unsupported int             `json:"unsupported_on_this_platform"`
	Percent     float64         `json:"percent"`
	Modules     []ModuleSummary `json:"modules"`
}

// ModuleSummary is per-module progress.
type ModuleSummary struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Total     int    `json:"total"`
	Completed int    `json:"completed"`
}

// Summarize computes progress. Exercises unsupported on goos are excluded
// from the denominator rather than counted as passed.
func (a *App) Summarize(st *progress.State, goos string) Summary {
	var s Summary
	mods := map[string]*ModuleSummary{}
	for _, m := range a.Catalog.Manifest.Modules {
		s.Modules = append(s.Modules, ModuleSummary{ID: m.ID, Title: m.Title})
	}
	for i := range s.Modules {
		mods[s.Modules[i].ID] = &s.Modules[i]
	}
	for _, v := range a.Views(st) {
		if goos != "" && !v.Exercise.SupportsOS(goos) {
			s.Unsupported++
			continue
		}
		s.Total++
		m := mods[v.Exercise.Module]
		m.Total++
		switch v.Status {
		case DisplayCompleted:
			s.Completed++
			m.Completed++
		case DisplayModified:
			s.Modified++
		case DisplaySkipped:
			s.Skipped++
		}
	}
	if s.Total > 0 {
		s.Percent = float64(s.Completed) * 100 / float64(s.Total)
	}
	// Hide modules that have no exercises yet.
	filtered := s.Modules[:0]
	for _, m := range s.Modules {
		if m.Total > 0 {
			filtered = append(filtered, m)
		}
	}
	s.Modules = filtered
	return s
}

// IsCanceled reports whether err stems from context cancellation.
func IsCanceled(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

// Hint is one revealed hint.
type Hint struct {
	Level int    `json:"level"`
	Total int    `json:"total"`
	Text  string `json:"text"`
	// AllRevealed is true when every hint has now been shown.
	AllRevealed bool `json:"all_revealed"`
}

// RevealHint shows hint level (1-based), or the next unrevealed hint when
// level is 0. Once every hint has been revealed, the last one is shown again.
// The highest revealed level is persisted.
func (a *App) RevealHint(ex *catalog.Exercise, level int) (Hint, error) {
	total := len(ex.Hints)
	if level < 0 || level > total {
		return Hint{}, fmt.Errorf("%s has %d hints; --level must be between 1 and %d", ex.ID, total, total)
	}
	var h Hint
	_, err := a.Store.Update(func(s *progress.State) error {
		e := s.Entry(ex.ID)
		if level == 0 {
			level = min(e.HintsRevealed+1, total)
		}
		e.HintsRevealed = max(e.HintsRevealed, level)
		if e.Status == progress.NotStarted {
			e.Status = progress.InProgress
		}
		h = Hint{Level: level, Total: total, Text: ex.Hints[level-1], AllRevealed: e.HintsRevealed >= total}
		return nil
	})
	return h, err
}

// RevealedHints returns the hints already revealed for ex.
func (a *App) RevealedHints(ex *catalog.Exercise, st *progress.State) []string {
	n := min(st.Get(ex.ID).HintsRevealed, len(ex.Hints))
	return ex.Hints[:n]
}

// MarkSolutionViewed records that the learner looked at the solution. It
// does not affect completion.
func (a *App) MarkSolutionViewed(ex *catalog.Exercise) error {
	_, err := a.Store.Update(func(s *progress.State) error {
		s.Entry(ex.ID).SolutionViewed = true
		return nil
	})
	return err
}
