// Mutant: unknown names return an empty circle instead of nil.
package find

import "math"

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }

func (c *Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

type Square struct{ Side float64 }

func (s *Square) Area() float64 { return s.Side * s.Side }

func Find(name string) Shape {
	switch name {
	case "circle":
		return &Circle{Radius: 1}
	case "square":
		return &Square{Side: 1}
	}
	return &Circle{}
}
