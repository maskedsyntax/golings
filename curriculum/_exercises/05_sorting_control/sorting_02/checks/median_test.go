package median

import (
	"slices"
	"testing"
)

func TestMedian(t *testing.T) {
	for _, tt := range []struct {
		nums []float64
		want float64
	}{
		{[]float64{3, 1, 2}, 2},
		{[]float64{4, 1, 3, 2}, 2.5},
		{[]float64{7}, 7},
		{[]float64{10, -5}, 2.5},
		{[]float64{9, 9, 1, 9, 2}, 9},
	} {
		in := slices.Clone(tt.nums)
		got, ok := Median(in)
		if !ok || got != tt.want {
			t.Errorf("Median(%v) = %v, %v; want %v, true", tt.nums, got, ok, tt.want)
		}
		if !slices.Equal(in, tt.nums) {
			t.Errorf("Median reordered its input to %v", in)
		}
	}
}

func TestMedianEmpty(t *testing.T) {
	if _, ok := Median(nil); ok {
		t.Fatal("Median(nil) reported ok = true")
	}
}
