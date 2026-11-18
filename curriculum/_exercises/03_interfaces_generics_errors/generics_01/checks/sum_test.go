package sum

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum([]int{1, 2, 3}) = %v, want 6", got)
	}
	if got := Sum([]int64{1 << 40, 1}); got != 1<<40+1 {
		t.Errorf("Sum([]int64) = %v", got)
	}
	if got := Sum([]float64{0.5, 0.25}); got != 0.75 {
		t.Errorf("Sum([]float64{0.5, 0.25}) = %v, want 0.75", got)
	}
	if got := Sum([]Cents{150, 50}); got != Cents(200) {
		t.Errorf("Sum([]Cents{150, 50}) = %v, want 200", got)
	}
	if got := Sum[int](nil); got != 0 {
		t.Errorf("Sum[int](nil) = %v, want 0", got)
	}
}
