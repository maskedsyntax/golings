package check

import (
	"errors"
	"slices"
	"testing"
)

func TestCheckValid(t *testing.T) {
	for _, words := range [][]string{{"go", "zig"}, nil, {"123", ""}} {
		if err := Check(words); err != nil {
			t.Errorf("Check(%q) = %#v, want a nil error", words, err)
		}
	}
}

func TestCheckInvalid(t *testing.T) {
	err := Check([]string{"Go", "rust", "ZIG"})
	var ce *CheckError
	if !errors.As(err, &ce) {
		t.Fatalf("Check returned %v (%T), want a *CheckError", err, err)
	}
	if !slices.Equal(ce.Problems, []string{"Go", "ZIG"}) {
		t.Errorf("Problems = %q, want [Go ZIG]", ce.Problems)
	}
	if err.Error() != "not lowercase: Go, ZIG" {
		t.Errorf("Error() = %q", err.Error())
	}
}
