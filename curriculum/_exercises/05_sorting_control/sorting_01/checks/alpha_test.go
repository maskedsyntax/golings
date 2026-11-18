package alpha

import (
	"slices"
	"testing"
)

func TestAlphabetize(t *testing.T) {
	words := []string{"pear", "apple", "fig", "banana"}
	got := Alphabetize(words)
	if want := []string{"apple", "banana", "fig", "pear"}; !slices.Equal(got, want) {
		t.Fatalf("Alphabetize = %q, want %q", got, want)
	}
	if want := []string{"pear", "apple", "fig", "banana"}; !slices.Equal(words, want) {
		t.Fatalf("the caller's slice was changed to %q; it should still be %q", words, want)
	}
}

func TestAlphabetizeEmpty(t *testing.T) {
	if got := Alphabetize(nil); len(got) != 0 {
		t.Fatalf("Alphabetize(nil) = %q", got)
	}
}
