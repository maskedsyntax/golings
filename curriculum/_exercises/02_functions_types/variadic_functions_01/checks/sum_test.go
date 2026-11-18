package sum

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}
	if got := Sum(1, 2, 3, 4); got != 10 {
		t.Errorf("Sum(1, 2, 3, 4) = %d, want 10", got)
	}
}

func TestAverage(t *testing.T) {
	for _, tt := range []struct {
		xs   []int
		want int
	}{{[]int{2, 4, 9}, 5}, {[]int{10}, 10}, {nil, 0}, {[]int{-3, 3}, 0}} {
		if got := Average(tt.xs); got != tt.want {
			t.Errorf("Average(%v) = %d, want %d", tt.xs, got, tt.want)
		}
	}
}

func TestTotal(t *testing.T) {
	if got := Total(1, 2, 3); got != 6 {
		t.Errorf("Total(1, 2, 3) = %d, want 6", got)
	}
}
