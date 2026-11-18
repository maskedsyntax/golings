package curve

import (
	"slices"
	"testing"
)

func TestCurve(t *testing.T) {
	s := []Student{{"Ana", 70}, {"Ben", 95}, {"Cy", 100}, {"Di", 0}}
	Curve(s, 10)
	want := []Student{{"Ana", 80}, {"Ben", 100}, {"Cy", 100}, {"Di", 10}}
	if !slices.Equal(s, want) {
		t.Errorf("after Curve(s, 10): %+v, want %+v", s, want)
	}
}

func TestCurveZeroBonus(t *testing.T) {
	s := []Student{{"Ana", 42}}
	Curve(s, 0)
	if s[0].Score != 42 {
		t.Errorf("Curve(s, 0) changed the score to %d", s[0].Score)
	}
	Curve(nil, 5) // must not panic
}
