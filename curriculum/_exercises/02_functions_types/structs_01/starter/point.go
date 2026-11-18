// structs_01: Fields and struct literals
//
// This file does not compile. Fix Point and NewPoint.
//
//	p := NewPoint(1, 2) // p.X == 1, p.Y == 2
//	Manhattan(Point{0, 0}, Point{3, -4}) == 7
//
// Reference: https://gobyexample.com/structs
package point

// Point is a location on a grid.
type Point struct {
	X int
}

// NewPoint returns a pointer to a new Point at (x, y).
func NewPoint(x, y int) *Point {
	p := Point{X: x, Y: y}
	return p
}

// Manhattan returns the grid distance between a and b.
func Manhattan(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
