package callall

import (
	"slices"
	"testing"
	"time"
)

func TestCallAll(t *testing.T) {
	hang := make(chan struct{})
	defer close(hang)
	hung := func() int { <-hang; return -1 }
	val := func(v int) func() int { return func() int { return v } }

	calls := []func() int{val(1), hung, val(3), hung, hung, val(6), hung}
	start := time.Now()
	got := CallAll(calls, 200*time.Millisecond)
	elapsed := time.Since(start)

	want := []Result{{1, false}, {0, true}, {3, false}, {0, true}, {0, true}, {6, false}, {0, true}}
	if !slices.Equal(got, want) {
		t.Fatalf("CallAll =\n  %v\nwant\n  %v", got, want)
	}
	if elapsed > 500*time.Millisecond {
		t.Fatalf("4 hung calls with a 200ms timeout took %v; run them concurrently so it takes about 200ms", elapsed)
	}
}
