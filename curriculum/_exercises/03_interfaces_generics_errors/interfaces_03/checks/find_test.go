package find

import (
	"math"
	"testing"
)

func TestFindKnown(t *testing.T) {
	if s := Find("circle"); s == nil || math.Abs(s.Area()-math.Pi) > 1e-9 {
		t.Errorf("Find(\"circle\") = %#v, want a unit circle", s)
	}
	if s := Find("square"); s == nil || s.Area() != 1 {
		t.Errorf("Find(\"square\") = %#v, want a unit square", s)
	}
}

func TestFindUnknownIsNil(t *testing.T) {
	for _, name := range []string{"nope", "", "Circle"} {
		if s := Find(name); s != nil {
			t.Errorf("Find(%q) = %#v, which is a non-nil Shape; want a nil interface", name, s)
		}
	}
}
