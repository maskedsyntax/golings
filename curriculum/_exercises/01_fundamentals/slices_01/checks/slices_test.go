package slices

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"three", []int{1, 2, 3}, []int{3, 2, 1}},
		{"even length", []int{4, 5, 6, 7}, []int{7, 6, 5, 4}},
		{"single", []int{42}, []int{42}},
		{"negatives and duplicates", []int{-1, 0, -1, 9}, []int{9, -1, 0, -1}},
		{"empty", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse(%v) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}

func TestReverseNilIsEmptyNonNil(t *testing.T) {
	got := Reverse(nil)
	if got == nil || len(got) != 0 {
		t.Errorf("Reverse(nil) = %#v, want an empty non-nil slice ([]int{})", got)
	}
}

func TestReverseDoesNotMutateInput(t *testing.T) {
	in := []int{1, 2, 3, 4}
	Reverse(in)
	if want := []int{1, 2, 3, 4}; !reflect.DeepEqual(in, want) {
		t.Errorf("input changed to %v after Reverse, want it untouched (%v)", in, want)
	}
}

func TestReverseReturnsNewBackingArray(t *testing.T) {
	in := []int{1, 2, 3}
	out := Reverse(in)
	if len(out) > 0 {
		out[0] = 100
		if in[2] == 100 {
			t.Errorf("writing to the result changed the input; Reverse must return a new slice")
		}
	}
}
