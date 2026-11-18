// interfaces_03: A nil that isn't nil
//
// Find("nope") != nil, so callers crash. Fix Find.
//
//	Find("circle") != nil
//	Find("nope") == nil
//
// Reference: https://gobyexample.com/interfaces
package find

import "math"

// Shape is anything with an area.
type Shape interface {
	Area() float64
}

// Circle is a circle.
type Circle struct{ Radius float64 }

func (c *Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// Square is a square.
type Square struct{ Side float64 }

func (s *Square) Area() float64 { return s.Side * s.Side }

// Find returns a unit shape with the given name, or nil if the name is unknown.
func Find(name string) Shape {
	var c *Circle
	var s *Square
	switch name {
	case "circle":
		c = &Circle{Radius: 1}
	case "square":
		s = &Square{Side: 1}
		return s
	}
	return c
}
