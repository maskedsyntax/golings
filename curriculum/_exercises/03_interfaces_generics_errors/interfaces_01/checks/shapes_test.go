package shapes

import (
	"math"
	"testing"
)

var _ Shape = Circle{} // Circle must satisfy Shape

func near(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestCirclePerimeter(t *testing.T) {
	if got := (Circle{Radius: 2}).Perimeter(); !near(got, 4*math.Pi) {
		t.Errorf("Circle{2}.Perimeter() = %v, want 4π (%v)", got, 4*math.Pi)
	}
}

func TestTotalPerimeter(t *testing.T) {
	shapes := []Shape{Rect{Width: 1, Height: 2}, Circle{Radius: 1}, Rect{Width: 3, Height: 3}}
	if got, want := TotalPerimeter(shapes), 6+2*math.Pi+12; !near(got, want) {
		t.Errorf("TotalPerimeter = %v, want %v", got, want)
	}
	if got := TotalPerimeter(nil); got != 0 {
		t.Errorf("TotalPerimeter(nil) = %v, want 0", got)
	}
}
