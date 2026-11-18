// json_04: Decoding by kind (reference solution)
package shapes

import (
	"encoding/json"
	"fmt"
	"math"
)

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
	var raw []struct {
		Kind string          `json:"kind"`
		Spec json.RawMessage `json:"spec"` // decoded below, once we know the kind
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	shapes := make([]Shape, 0, len(raw))
	for i, r := range raw {
		var s Shape
		var err error
		switch r.Kind {
		case "circle":
			var c Circle
			err = json.Unmarshal(r.Spec, &c)
			s = c
		case "rect":
			var rc Rect
			err = json.Unmarshal(r.Spec, &rc)
			s = rc
		default:
			return nil, fmt.Errorf("shape %d: unknown shape kind %q", i, r.Kind)
		}
		if err != nil {
			return nil, fmt.Errorf("shape %d (%s): %w", i, r.Kind, err)
		}
		shapes = append(shapes, s)
	}
	return shapes, nil
}
