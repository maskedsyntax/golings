package tags

import (
	"slices"
	"testing"
)

func TestHashTags(t *testing.T) {
	for _, tt := range []struct {
		text string
		want []string
	}{
		{"Loving #Go and #rust_lang, #go again; # alone", []string{"go", "rust_lang"}},
		{"#a #b #A #c #B", []string{"a", "b", "c"}},
		{"#2024 recap", []string{"2024"}},
		{"no tags here", nil},
		{"", nil},
	} {
		got := HashTags(tt.text)
		if len(got) != len(tt.want) || !slices.Equal(got, tt.want) {
			t.Errorf("HashTags(%q) = %q, want %q", tt.text, got, tt.want)
		}
	}
}
