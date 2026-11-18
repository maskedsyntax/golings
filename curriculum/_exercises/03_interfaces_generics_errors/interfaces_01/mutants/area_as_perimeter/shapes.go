// Mutant: Circle.Perimeter returns the area.
package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct{ Width, Height float64 }

func (r Rect) Area() float64      { return r.Width * r.Height }
func (r Rect) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return c.Area() }

func TotalPerimeter(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Perimeter()
	}
	return total
}
