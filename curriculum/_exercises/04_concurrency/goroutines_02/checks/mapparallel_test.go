package mapparallel

import (
	"slices"
	"strings"
	"testing"
	"time"
)

func TestMapParallel(t *testing.T) {
	slowUpper := func(s string) string {
		time.Sleep(10 * time.Millisecond)
		return strings.ToUpper(s)
	}
	got := MapParallel([]string{"a", "b", "c", "d"}, slowUpper)
	if want := []string{"A", "B", "C", "D"}; !slices.Equal(got, want) {
		t.Fatalf("MapParallel = %q, want %q", got, want)
	}
}

func TestMapParallelIsConcurrent(t *testing.T) {
	items := make([]string, 20)
	start := time.Now()
	MapParallel(items, func(s string) string { time.Sleep(50 * time.Millisecond); return s })
	if d := time.Since(start); d > 600*time.Millisecond {
		t.Fatalf("20 calls of 50ms took %v; they should run at the same time", d)
	}
}

func TestMapParallelEmpty(t *testing.T) {
	if got := MapParallel(nil, strings.ToUpper); len(got) != 0 {
		t.Fatalf("MapParallel(nil) = %q", got)
	}
}
