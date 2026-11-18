package dice

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestRollRange(t *testing.T) {
	rolls := Roll(rand.New(rand.NewPCG(1, 2)), 600)
	if len(rolls) != 600 {
		t.Fatalf("Roll returned %d rolls, want 600", len(rolls))
	}
	seen := make(map[int]int)
	for _, v := range rolls {
		if v < 1 || v > 6 {
			t.Fatalf("Roll returned %d; a die shows 1 to 6", v)
		}
		seen[v]++
	}
	for face := 1; face <= 6; face++ {
		if seen[face] == 0 {
			t.Errorf("600 rolls never showed %d", face)
		}
	}
}

func TestRollReproducible(t *testing.T) {
	a := Roll(rand.New(rand.NewPCG(7, 7)), 20)
	b := Roll(rand.New(rand.NewPCG(7, 7)), 20)
	if !slices.Equal(a, b) {
		t.Fatalf("the same seed gave different rolls:\n%v\n%v", a, b)
	}
	if got := Roll(rand.New(rand.NewPCG(1, 1)), 0); len(got) != 0 {
		t.Fatalf("Roll(r, 0) = %v, want no rolls", got)
	}
}
