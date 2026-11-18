package run

import (
	"slices"
	"testing"
)

func TestRun(t *testing.T) {
	got := Run([]string{"a", "b", "c"})
	want := []string{"run a", "run b", "run c", "cleanup after c"}
	if !slices.Equal(got, want) {
		t.Fatalf("Run = %q\nwant  %q", got, want)
	}
}

func TestRunOneStep(t *testing.T) {
	got := Run([]string{"only"})
	if want := []string{"run only", "cleanup after only"}; !slices.Equal(got, want) {
		t.Fatalf("Run = %q, want %q", got, want)
	}
}
