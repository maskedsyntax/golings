package multipliers

import "testing"

func TestMultipliers(t *testing.T) {
	factors := []int{2, 3, 10, -1}
	fns := Multipliers(factors)
	if len(fns) != len(factors) {
		t.Fatalf("got %d functions, want %d", len(fns), len(factors))
	}
	for i, f := range factors {
		if got := fns[i](5); got != 5*f {
			t.Errorf("fns[%d](5) = %d, want %d", i, got, 5*f)
		}
	}
}

func TestMultipliersIgnoreLaterChanges(t *testing.T) {
	factors := []int{4, 7}
	fns := Multipliers(factors)
	factors[0] = 100
	if got := fns[0](1); got != 4 {
		t.Errorf("after changing factors[0], fns[0](1) = %d, want 4 (the factor when Multipliers was called)", got)
	}
}

func TestMultipliersEmpty(t *testing.T) {
	if fns := Multipliers(nil); len(fns) != 0 {
		t.Errorf("Multipliers(nil) returned %d functions, want 0", len(fns))
	}
}
