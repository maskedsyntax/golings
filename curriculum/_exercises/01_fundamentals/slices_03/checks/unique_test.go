package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		in, want []int
	}{
		{"repeated", []int{1, 2, 1, 3, 2}, []int{1, 2, 3}},
		{"first-seen order", []int{3, 1, 3, 2, 1}, []int{3, 1, 2}},
		{"negatives", []int{-1, 0, -1}, []int{-1, 0}},
		{"all same", []int{4, 4, 4}, []int{4}},
		{"empty", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Unique(%v) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}

func TestUniqueNil(t *testing.T) {
	if got := Unique(nil); got == nil || len(got) != 0 {
		t.Errorf("Unique(nil) = %#v, want []int{}", got)
	}
}

func TestUniqueKeepsInput(t *testing.T) {
	in := []int{2, 1, 2}
	Unique(in)
	if !reflect.DeepEqual(in, []int{2, 1, 2}) {
		t.Errorf("input changed to %v", in)
	}
}
