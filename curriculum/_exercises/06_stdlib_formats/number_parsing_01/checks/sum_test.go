package sum

import (
	"errors"
	"strconv"
	"testing"
)

func TestSum(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want int
	}{
		{"1 2 3", 6},
		{"10 -4", 6},
		{"  7  ", 7},
		{"", 0},
	} {
		got, err := Sum(tt.in)
		if err != nil || got != tt.want {
			t.Errorf("Sum(%q) = %d, %v; want %d, nil", tt.in, got, err, tt.want)
		}
	}
}

func TestSumError(t *testing.T) {
	_, err := Sum("1 two 3")
	if !errors.Is(err, strconv.ErrSyntax) {
		t.Fatalf("Sum(\"1 two 3\") error = %v, want the strconv syntax error", err)
	}
}
