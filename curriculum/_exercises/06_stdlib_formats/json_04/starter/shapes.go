// json_04: Decoding by kind
//
// Implement DecodeShapes. See README.md.
//
// Reference: https://gobyexample.com/json
package shapes

import "math"

// Shape has an area.
type Shape interface {
	Area() float64
}

// Circle has a radius.
type Circle struct {
	R float64 `json:"r"`
}

// Area implements Shape.
func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

// Rect has a width and height.
type Rect struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

// Area implements Shape.
func (r Rect) Area() float64 { return r.W * r.H }

// DecodeShapes decodes a JSON array of {"kind": ..., "spec": {...}} objects.
func DecodeShapes(data []byte) ([]Shape, error) {
	// TODO
	return nil, nil
}
