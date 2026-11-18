package pipeline

import (
	"iter"
	"slices"
	"testing"
)

func isOdd(n int) bool { return n%2 == 1 }

func TestPipeline(t *testing.T) {
	if got := slices.Collect(Take(Filter(Naturals(), isOdd), 3)); !slices.Equal(got, []int{1, 3, 5}) {
		t.Errorf("Take(Filter(Naturals, odd), 3) = %v, want [1 3 5]", got)
	}
	if got := slices.Collect(Take(Naturals(), 4)); !slices.Equal(got, []int{0, 1, 2, 3}) {
		t.Errorf("Take(Naturals, 4) = %v, want [0 1 2 3]", got)
	}
	words := slices.Values([]string{"go", "", "iter", ""})
	if got := slices.Collect(Filter(words, func(s string) bool { return s != "" })); !slices.Equal(got, []string{"go", "iter"}) {
		t.Errorf("Filter(non-empty) = %q", got)
	}
	if got := slices.Collect(Take(slices.Values([]int{7, 8}), 5)); !slices.Equal(got, []int{7, 8}) {
		t.Errorf("Take of a short sequence = %v, want [7 8]", got)
	}
}

// counting wraps a sequence and counts how many values were pulled from it.
func counting(seq iter.Seq[int], pulled *int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for v := range seq {
			*pulled++
			if !yield(v) {
				return
			}
		}
	}
}

func TestTakePullsOnlyWhatItNeeds(t *testing.T) {
	pulled := 0
	slices.Collect(Take(counting(Naturals(), &pulled), 3))
	if pulled != 3 {
		t.Errorf("Take(seq, 3) pulled %d values from seq, want 3", pulled)
	}
	pulled = 0
	for range Take(counting(Naturals(), &pulled), 0) {
		t.Fatal("Take(seq, 0) yielded a value")
	}
	if pulled != 0 {
		t.Errorf("Take(seq, 0) pulled %d values, want 0", pulled)
	}
}

func TestStopOnBreak(t *testing.T) {
	var seen []int
	for v := range Filter(Take(Naturals(), 100), isOdd) {
		seen = append(seen, v)
		if len(seen) == 2 {
			break
		}
	}
	if !slices.Equal(seen, []int{1, 3}) {
		t.Errorf("saw %v, want [1 3]", seen)
	}
}
