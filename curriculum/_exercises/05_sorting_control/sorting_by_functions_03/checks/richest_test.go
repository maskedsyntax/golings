package richest

import (
	"math"
	"slices"
	"testing"
)

func owners(accts []Account) []string {
	var out []string
	for _, a := range accts {
		out = append(out, a.Owner)
	}
	return out
}

func TestRichestSmall(t *testing.T) {
	accts := []Account{{"a", 100}, {"b", 300}, {"c", -50}, {"d", 200}}
	Richest(accts)
	if got, want := owners(accts), []string{"b", "d", "a", "c"}; !slices.Equal(got, want) {
		t.Fatalf("Richest = %v, want %v", got, want)
	}
}

func TestRichestHugeBalances(t *testing.T) {
	accts := []Account{
		{"overdrawn", math.MinInt64 + 1},
		{"whale", math.MaxInt64 - 1},
		{"zero", 0},
	}
	Richest(accts)
	if got, want := owners(accts), []string{"whale", "zero", "overdrawn"}; !slices.Equal(got, want) {
		t.Fatalf("Richest = %v, want %v", got, want)
	}
}
