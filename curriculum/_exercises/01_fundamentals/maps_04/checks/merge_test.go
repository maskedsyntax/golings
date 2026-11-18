package merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 3, "z": -1}
	got := Merge(a, b)
	if want := map[string]int{"x": 1, "y": 5, "z": -1}; !reflect.DeepEqual(got, want) {
		t.Errorf("Merge(%v, %v) = %v, want %v", a, b, got, want)
	}
	if !reflect.DeepEqual(a, map[string]int{"x": 1, "y": 2}) || !reflect.DeepEqual(b, map[string]int{"y": 3, "z": -1}) {
		t.Errorf("Merge modified its inputs: a = %v, b = %v", a, b)
	}
}

func TestMergeNilInputs(t *testing.T) {
	got := Merge(nil, nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Merge(nil, nil) = %#v, want an empty non-nil map", got)
	}
	if got := Merge(nil, map[string]int{"k": 1}); !reflect.DeepEqual(got, map[string]int{"k": 1}) {
		t.Errorf("Merge(nil, {k:1}) = %v", got)
	}
}

func TestMergeResultIsIndependent(t *testing.T) {
	a := map[string]int{"x": 1}
	out := Merge(a, nil)
	out["x"] = 100
	if a["x"] != 1 {
		t.Error("writing to the result changed the first input")
	}
}
