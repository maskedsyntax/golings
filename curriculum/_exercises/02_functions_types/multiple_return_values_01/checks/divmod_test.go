package divmod

import "testing"

func TestDivMod(t *testing.T) {
	for _, tt := range []struct{ a, b, q, r int }{{7, 2, 3, 1}, {9, 3, 3, 0}, {-7, 2, -3, -1}, {1, 5, 0, 1}} {
		q, r := DivMod(tt.a, tt.b)
		if q != tt.q || r != tt.r {
			t.Errorf("DivMod(%d, %d) = %d, %d, want %d, %d", tt.a, tt.b, q, r, tt.q, tt.r)
		}
	}
}

func TestQuotient(t *testing.T) {
	if got := Quotient(17, 5); got != 3 {
		t.Errorf("Quotient(17, 5) = %d, want 3", got)
	}
}
