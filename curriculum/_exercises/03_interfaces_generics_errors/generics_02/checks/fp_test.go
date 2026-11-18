package fp

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	words := []string{"go", "is", "fun"}
	if got := Map(words, strings.ToUpper); !slices.Equal(got, []string{"GO", "IS", "FUN"}) {
		t.Errorf("Map(words, ToUpper) = %q", got)
	}
	if got := Map([]int{1, 22, 333}, strconv.Itoa); !slices.Equal(got, []string{"1", "22", "333"}) {
		t.Errorf("Map(ints, Itoa) = %q", got)
	}
	if got := Map(words, func(s string) int { return len(s) }); !slices.Equal(got, []int{2, 2, 3}) {
		t.Errorf("Map(words, len) = %v", got)
	}
	if !slices.Equal(words, []string{"go", "is", "fun"}) {
		t.Errorf("Map modified its input: %q", words)
	}
}

func TestFilter(t *testing.T) {
	nums := []int{5, -2, 8, 0, -7}
	if got := Filter(nums, func(n int) bool { return n > 0 }); !slices.Equal(got, []int{5, 8}) {
		t.Errorf("Filter(positive) = %v, want [5 8]", got)
	}
	if !slices.Equal(nums, []int{5, -2, 8, 0, -7}) {
		t.Errorf("Filter modified its input: %v", nums)
	}
	if got := Filter([]string{"a", "", "b"}, func(s string) bool { return s != "" }); !slices.Equal(got, []string{"a", "b"}) {
		t.Errorf("Filter(non-empty) = %q", got)
	}
}

func TestReduce(t *testing.T) {
	sum := func(acc, x int) int { return acc + x }
	if got := Reduce([]int{1, 2, 3}, 10, sum); got != 16 {
		t.Errorf("Reduce([1 2 3], 10, +) = %d, want 16", got)
	}
	if got := Reduce(nil, 42, sum); got != 42 {
		t.Errorf("Reduce(nil, 42, +) = %d, want 42", got)
	}
	join := func(acc string, n int) string { return acc + strconv.Itoa(n) }
	if got := Reduce([]int{1, 2, 3}, ">", join); got != ">123" {
		t.Errorf("Reduce left to right = %q, want \">123\"", got)
	}
}
