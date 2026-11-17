package maps

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  map[string]int
	}{
		{"repeats", []string{"go", "gopher", "go"}, map[string]int{"go": 2, "gopher": 1}},
		{"case sensitive", []string{"go", "Go", "go"}, map[string]int{"go": 2, "Go": 1}},
		{"unicode", []string{"héllo", "héllo"}, map[string]int{"héllo": 2}},
		{"empty", []string{}, map[string]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordCount(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordCount(%#v) = %#v, want %#v", tt.words, got, tt.want)
			}
		})
	}
}

func TestWordCountNilInputGivesWritableMap(t *testing.T) {
	got := WordCount(nil)
	if got == nil {
		t.Fatal("WordCount(nil) returned a nil map; callers expect a non-nil map they can write to")
	}
	if len(got) != 0 {
		t.Errorf("WordCount(nil) = %v, want an empty map", got)
	}
}

func TestHas(t *testing.T) {
	m := map[string]int{"zero": 0, "one": 1}
	tests := []struct {
		key  string
		want bool
	}{
		{"one", true},
		{"zero", true},
		{"missing", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := Has(m, tt.key); got != tt.want {
				t.Errorf("Has(%v, %q) = %v, want %v", m, tt.key, got, tt.want)
			}
		})
	}
}

func TestHasNilMap(t *testing.T) {
	if Has(nil, "anything") {
		t.Error("Has(nil, \"anything\") = true, want false")
	}
}
