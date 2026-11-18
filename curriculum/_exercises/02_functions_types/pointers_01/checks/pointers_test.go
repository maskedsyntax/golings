package pointers

import "testing"

func TestDouble(t *testing.T) {
	x := 4
	Double(&x)
	if x != 8 {
		t.Errorf("after Double(&x) with x = 4, x = %d, want 8", x)
	}
	Double(&x)
	if x != 16 {
		t.Errorf("after a second Double, x = %d, want 16", x)
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	pa, pb := &a, &b
	Swap(pa, pb)
	if a != 2 || b != 1 {
		t.Errorf("after Swap(&a, &b) with a = 1, b = 2: a = %d, b = %d, want a = 2, b = 1", a, b)
	}
	if pa != &a || pb != &b {
		t.Error("Swap must swap the values, not the caller's pointers")
	}
}
