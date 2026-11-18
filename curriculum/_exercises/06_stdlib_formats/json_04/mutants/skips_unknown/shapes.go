// Mutant: silently skips unknown kinds.
package shapes

import (
	"encoding/json"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	R float64 `json:"r"`
}

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

type Rect struct {
	W float64 `json:"w"`
	H float64 `json:"h"`
}

func (r Rect) Area() float64 { return r.W * r.H }

func DecodeShapes(data []byte) ([]Shape, error) {
	var raw []struct {
		Kind string          `json:"kind"`
		Spec json.RawMessage `json:"spec"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	var shapes []Shape
	for _, r := range raw {
		switch r.Kind {
		case "circle":
			var c Circle
			if err := json.Unmarshal(r.Spec, &c); err != nil {
				return nil, err
			}
			shapes = append(shapes, c)
		case "rect":
			var rc Rect
			if err := json.Unmarshal(r.Spec, &rc); err != nil {
				return nil, err
			}
			shapes = append(shapes, rc)
		}
	}
	return shapes, nil
}
