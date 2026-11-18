package point

import "testing"

func TestNewPoint(t *testing.T) {
	p := NewPoint(1, 2)
	if p == nil || p.X != 1 || p.Y != 2 {
		t.Fatalf("NewPoint(1, 2) = %+v, want &{X:1 Y:2}", p)
	}
	if q := NewPoint(1, 2); q == p {
		t.Error("two calls to NewPoint returned the same pointer")
	}
}

func TestManhattan(t *testing.T) {
	for _, tt := range []struct {
		a, b Point
		want int
	}{
		{Point{0, 0}, Point{3, -4}, 7},
		{Point{X: 2, Y: 2}, Point{X: 2, Y: 2}, 0},
		{Point{Y: 5}, Point{X: -1}, 6},
	} {
		if got := Manhattan(tt.a, tt.b); got != tt.want {
			t.Errorf("Manhattan(%+v, %+v) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
