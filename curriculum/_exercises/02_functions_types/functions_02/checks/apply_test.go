package apply

import (
	"slices"
	"testing"
)

func double(n int) int    { return n * 2 }
func isEven(n int) bool   { return n%2 == 0 }
func negate(n int) int    { return -n }
func positive(n int) bool { return n > 0 }

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3}
	if got := Map(nums, double); !slices.Equal(got, []int{2, 4, 6}) {
		t.Errorf("Map(%v, double) = %v, want [2 4 6]", nums, got)
	}
	if got := Map([]int{5, -1}, negate); !slices.Equal(got, []int{-5, 1}) {
		t.Errorf("Map([5 -1], negate) = %v, want [-5 1]", got)
	}
	if got := Map(nil, double); len(got) != 0 {
		t.Errorf("Map(nil, double) = %v, want empty", got)
	}
	if !slices.Equal(nums, []int{1, 2, 3}) {
		t.Errorf("Map modified its input: nums = %v", nums)
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, -6}
	if got := Filter(nums, isEven); !slices.Equal(got, []int{2, 4, -6}) {
		t.Errorf("Filter(%v, isEven) = %v, want [2 4 -6]", nums, got)
	}
	if got := Filter(nums, positive); !slices.Equal(got, []int{1, 2, 3, 4}) {
		t.Errorf("Filter(%v, positive) = %v, want [1 2 3 4]", nums, got)
	}
	if got := Filter([]int{1, 3}, isEven); len(got) != 0 {
		t.Errorf("Filter([1 3], isEven) = %v, want empty", got)
	}
	if !slices.Equal(nums, []int{1, 2, 3, 4, -6}) {
		t.Errorf("Filter modified its input: nums = %v", nums)
	}
}
