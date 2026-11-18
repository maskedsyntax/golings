package group

import (
	"reflect"
	"testing"
)

func TestGroupByLength(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[int][]string
	}{
		{"basic", []string{"go", "is", "fun", "ok"}, map[int][]string{2: {"go", "is", "ok"}, 3: {"fun"}}},
		{"runes not bytes", []string{"héllo", "world"}, map[int][]string{5: {"héllo", "world"}}},
		{"empty word", []string{"", "a"}, map[int][]string{0: {""}, 1: {"a"}}},
		{"duplicates kept", []string{"hi", "hi"}, map[int][]string{2: {"hi", "hi"}}},
		{"no words", []string{}, map[int][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupByLength(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByLength(%q) = %v, want %v", tt.words, got, tt.want)
			}
		})
	}
}

func TestGroupByLengthNil(t *testing.T) {
	if got := GroupByLength(nil); got == nil {
		t.Error("GroupByLength(nil) returned a nil map")
	}
}
