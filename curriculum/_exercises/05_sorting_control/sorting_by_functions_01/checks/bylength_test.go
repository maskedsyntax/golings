package bylength

import (
	"slices"
	"testing"
)

func TestByLength(t *testing.T) {
	words := []string{"peach", "banana", "kiwi", "fig", "apple", "date"}
	ByLength(words)
	want := []string{"fig", "date", "kiwi", "apple", "peach", "banana"}
	if !slices.Equal(words, want) {
		t.Fatalf("ByLength = %q, want %q", words, want)
	}
}
