// structs_01: Fields and struct literals (reference solution)
package point

// Point is a location on a grid.
type Point struct {
	X, Y int
}

// NewPoint returns a pointer to a new Point at (x, y).
func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
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
