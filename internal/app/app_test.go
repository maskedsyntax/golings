package app_test

import (
	"path/filepath"
	"testing"

	"github.com/maskedsyntax/golings/curriculum"
	"github.com/maskedsyntax/golings/internal/app"
	"github.com/maskedsyntax/golings/internal/catalog"
	"github.com/maskedsyntax/golings/internal/progress"
	"github.com/maskedsyntax/golings/internal/workspace"
)

// On a platform an exercise doesn't support, it must never become the next
// exercise automatically, or the learner would be stuck on it.
func TestNextSkipsUnsupportedExercises(t *testing.T) {
	cat, err := catalog.Load(curriculum.FS)
	if err != nil {
		t.Fatal(err)
	}
	ws, err := workspace.Init(filepath.Join(t.TempDir(), "ws"), cat, workspace.InitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	a := app.New(cat, ws)
	a.GOOS = "windows"

	// Pretend everything before signals_02 (Linux and macOS only) is done.
	st := &progress.State{Exercises: map[string]*progress.ExerciseState{}}
	var target *catalog.Exercise
	for _, ex := range cat.Exercises() {
		if ex.ID == "signals_02" {
			target = ex
			break
		}
		st.Exercises[ex.ID] = &progress.ExerciseState{Status: progress.Completed}
	}
	if target == nil || target.SupportsOS("windows") {
		t.Fatal("test setup: signals_02 should exist and exclude windows")
	}
	st.CurrentExerciseID = "signals_02"

	if got := a.NextUnfinished(st); got == nil || got.ID == "signals_02" {
		t.Fatalf("NextUnfinished on windows = %v; want the exercise after signals_02", got)
	}
	if got := a.Current(st); got == nil || got.ID == "signals_02" {
		t.Fatalf("Current on windows = %v; want to move past signals_02", got)
	}
	a.GOOS = "linux"
	if got := a.Current(st); got == nil || got.ID != "signals_02" {
		t.Fatalf("Current on linux = %v; want signals_02", got)
	}
}
