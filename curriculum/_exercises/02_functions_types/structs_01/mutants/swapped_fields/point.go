// Mutant: NewPoint swaps x and y.
package point

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point { return &Point{X: y, Y: x} }

func Manhattan(a, b Point) int { return abs(a.X-b.X) + abs(a.Y-b.Y) }

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
