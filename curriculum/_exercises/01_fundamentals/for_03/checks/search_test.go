package search

import "testing"

func TestPositiveSum(t *testing.T) {
	for _, tt := range []struct {
		in   []int
		want int
	}{
		{[]int{1, -2, 3, 0, 4}, 8}, {[]int{-1, -2}, 0}, {nil, 0},
	} {
		if got := PositiveSum(tt.in); got != tt.want {
			t.Errorf("PositiveSum(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestFirstPair(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		i, j   int
		found  bool
	}{
		{"basic", []int{3, 1, 4, 2}, 5, 0, 3, true},
		{"several pairs: smallest i wins", []int{1, 4, 2, 3}, 5, 0, 1, true},
		{"same i, smallest j wins", []int{2, 3, 3}, 5, 0, 1, true},
		{"no reuse of one element", []int{5, 1}, 10, 0, 0, false},
		{"negatives", []int{-3, 7, 10}, 7, 0, 2, true},
		{"empty", nil, 1, 0, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, j, found := FirstPair(tt.nums, tt.target)
			if i != tt.i || j != tt.j || found != tt.found {
				t.Errorf("FirstPair(%v, %d) = (%d, %d, %v), want (%d, %d, %v)",
					tt.nums, tt.target, i, j, found, tt.i, tt.j, tt.found)
			}
		})
	}
}
