package shapes

import (
	"math"
	"strings"
	"testing"
)

func TestDecodeShapes(t *testing.T) {
	in := `[
		{"kind": "circle", "spec": {"r": 1}},
		{"kind": "rect", "spec": {"w": 2, "h": 3}},
		{"kind": "circle", "spec": {"r": 2}}
	]`
	got, err := DecodeShapes([]byte(in))
	if err != nil {
		t.Fatalf("DecodeShapes returned an error: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("DecodeShapes returned %d shapes, want 3", len(got))
	}
	if c, ok := got[0].(Circle); !ok || c.R != 1 {
		t.Errorf("shape 0 = %#v, want Circle{R: 1}", got[0])
	}
	if r, ok := got[1].(Rect); !ok || r.W != 2 || r.H != 3 {
		t.Errorf("shape 1 = %#v, want Rect{W: 2, H: 3}", got[1])
	}
	total := 0.0
	for _, s := range got {
		total += s.Area()
	}
	if want := 5*math.Pi + 6; math.Abs(total-want) > 1e-9 {
		t.Errorf("total area = %v, want %v", total, want)
	}
}

func TestDecodeShapesEmpty(t *testing.T) {
	got, err := DecodeShapes([]byte(`[]`))
	if err != nil || len(got) != 0 {
		t.Fatalf("DecodeShapes([]) = %v, %v; want no shapes and no error", got, err)
	}
}

func TestDecodeShapesErrors(t *testing.T) {
	_, err := DecodeShapes([]byte(`[{"kind": "hexagon", "spec": {}}]`))
	if err == nil {
		t.Fatal("an unknown kind should be an error")
	}
	if !strings.Contains(err.Error(), "hexagon") {
		t.Errorf("error %q should name the unknown kind", err)
	}
	for _, in := range []string{
		`[{"kind": "rect", "spec": {"w": "wide", "h": 1}}]`,
		`{"kind": "circle"}`,
		`[`,
	} {
		if _, err := DecodeShapes([]byte(in)); err == nil {
			t.Errorf("DecodeShapes(%s) returned no error", in)
		}
	}
}
