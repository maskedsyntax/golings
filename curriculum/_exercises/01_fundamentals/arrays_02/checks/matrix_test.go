package matrix

import "testing"

func TestTranspose(t *testing.T) {
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	if got := Transpose(m); got != want {
		t.Errorf("Transpose(%v) = %v, want %v", m, got, want)
	}
	if got := Transpose(Transpose(m)); got != m {
		t.Errorf("Transpose twice = %v, want the original %v", got, m)
	}
}

func TestEqual(t *testing.T) {
	a := [3][3]int{{1, 2, 3}}
	b := a
	if !Equal(a, b) {
		t.Errorf("Equal(%v, %v) = false, want true", a, b)
	}
	b[2][2] = 1
	if Equal(a, b) {
		t.Errorf("Equal(%v, %v) = true, want false", a, b)
	}
	if !Equal([3][3]int{}, [3][3]int{}) {
		t.Error("two zero matrices should be equal")
	}
}
