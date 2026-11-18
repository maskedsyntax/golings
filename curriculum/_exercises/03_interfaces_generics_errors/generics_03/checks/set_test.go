package set

import "testing"

func TestZeroValue(t *testing.T) {
	var s Set[string]
	if s.Len() != 0 || s.Has("x") {
		t.Fatal("a new set should be empty")
	}
	s.Add("go")
	s.Add("go")
	s.Add("rust")
	if s.Len() != 2 {
		t.Errorf("Len() = %d after adding go, go, rust; want 2", s.Len())
	}
	if !s.Has("go") || !s.Has("rust") || s.Has("zig") {
		t.Error("Has reports the wrong membership")
	}
}

type point struct{ X, Y int }

func TestStructElements(t *testing.T) {
	var s Set[point]
	s.Add(point{1, 2})
	s.Add(point{1, 2})
	if s.Len() != 1 || !s.Has(point{1, 2}) || s.Has(point{2, 1}) {
		t.Error("Set[point] doesn't treat equal structs as one element")
	}
}

func TestUnion(t *testing.T) {
	var a, b Set[int]
	for _, v := range []int{1, 2, 3} {
		a.Add(v)
	}
	for _, v := range []int{3, 4} {
		b.Add(v)
	}
	u := Union(&a, &b)
	if u.Len() != 4 {
		t.Errorf("Union has %d elements, want 4", u.Len())
	}
	for _, v := range []int{1, 2, 3, 4} {
		if !u.Has(v) {
			t.Errorf("Union is missing %d", v)
		}
	}
	u.Add(99)
	if a.Has(99) || b.Has(99) || a.Len() != 3 || b.Len() != 2 {
		t.Error("Union must return a new set; a or b changed")
	}
	var e Set[int]
	if Union(&e, &e).Len() != 0 {
		t.Error("Union of two empty sets should be empty")
	}
}
