// interfaces_01: Satisfying an interface
//
// The tests don't compile: Circle is not a Shape. Fix it.
//
//	TotalPerimeter([]Shape{Rect{1, 2}, Circle{1}}) == 6 + 2π
//
// Reference: https://gobyexample.com/interfaces
package shapes

import "math"

// Shape is anything with an area and a perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rect is a rectangle.
type Rect struct{ Width, Height float64 }

func (r Rect) Area() float64      { return r.Width * r.Height }
func (r Rect) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Circle is a circle.
type Circle struct{ Radius float64 }

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// TotalPerimeter adds up the perimeters of all shapes.
func TotalPerimeter(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Perimeter()
	}
	return total
}
