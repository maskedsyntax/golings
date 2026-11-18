package constants

import "testing"

func TestScale(t *testing.T) {
	for _, tt := range []struct{ in, want float64 }{{1.5, 3}, {0, 0}, {-2.25, -4.5}} {
		if got := Scale(tt.in); got != tt.want {
			t.Errorf("Scale(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestTB(t *testing.T) {
	var tb int64 = TB
	if tb != 1099511627776 {
		t.Errorf("TB = %d, want 1099511627776", tb)
	}
}

func TestFactorIsStillTwo(t *testing.T) {
	var n int = Factor
	if n != 2 {
		t.Errorf("Factor = %d, want 2", n)
	}
}
