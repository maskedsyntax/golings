package sample

import (
	"math/rand/v2"
	"slices"
	"testing"
)

var fruit = []string{"apple", "banana", "cherry", "date", "elder"}

func TestSampleShape(t *testing.T) {
	items := slices.Clone(fruit)
	r := rand.New(rand.NewPCG(3, 4))
	for k := 0; k <= len(items)+2; k++ {
		got := Sample(r, items, k)
		want := min(k, len(items))
		if len(got) != want {
			t.Fatalf("Sample(k=%d) returned %d items, want %d", k, len(got), want)
		}
		seen := map[string]bool{}
		for _, s := range got {
			if !slices.Contains(fruit, s) {
				t.Fatalf("Sample returned %q, which isn't one of the items", s)
			}
			if seen[s] {
				t.Fatalf("Sample(k=%d) = %q repeats %q", k, got, s)
			}
			seen[s] = true
		}
	}
	if !slices.Equal(items, fruit) {
		t.Fatalf("Sample reordered the caller's slice: %q", items)
	}
	if got := Sample(r, items, -1); len(got) != 0 {
		t.Fatalf("Sample(k=-1) = %q, want nothing", got)
	}
}

func TestSampleReproducible(t *testing.T) {
	a := Sample(rand.New(rand.NewPCG(9, 9)), fruit, 3)
	b := Sample(rand.New(rand.NewPCG(9, 9)), fruit, 3)
	if !slices.Equal(a, b) {
		t.Fatalf("the same seed gave %q and %q", a, b)
	}
}

func TestSampleIsFair(t *testing.T) {
	r := rand.New(rand.NewPCG(5, 6))
	counts := map[string]int{}
	const trials = 5000
	for range trials {
		counts[Sample(r, fruit, 1)[0]]++
	}
	for _, f := range fruit {
		// Expected 1000 each; 800..1200 is far outside normal variation.
		if c := counts[f]; c < 800 || c > 1200 {
			t.Errorf("%q was picked %d times out of %d; expected about %d", f, c, trials, trials/len(fruit))
		}
	}
}
