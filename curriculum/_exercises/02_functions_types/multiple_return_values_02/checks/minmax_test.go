package minmax

import "testing"

func TestMinMax(t *testing.T) {
	for _, tt := range []struct {
		nums   []int
		lo, hi int
	}{
		{[]int{3, 9, -2}, -2, 9},
		{[]int{4, 5, 6}, 4, 6},
		{[]int{-5, -1, -9}, -9, -1},
		{[]int{7}, 7, 7},
	} {
		lo, hi, ok := MinMax(tt.nums)
		if lo != tt.lo || hi != tt.hi || !ok {
			t.Errorf("MinMax(%v) = %d, %d, %v, want %d, %d, true", tt.nums, lo, hi, ok, tt.lo, tt.hi)
		}
	}
}

func TestMinMaxEmpty(t *testing.T) {
	for _, nums := range [][]int{nil, {}} {
		if lo, hi, ok := MinMax(nums); lo != 0 || hi != 0 || ok {
			t.Errorf("MinMax(%#v) = %d, %d, %v, want 0, 0, false", nums, lo, hi, ok)
		}
	}
}
