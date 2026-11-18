package kv

import (
	"maps"
	"testing"
)

func TestParseKV(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want map[string]string
	}{
		{"name=Ada; lang = Go ;expr=a=b;; novalue ; empty=",
			map[string]string{"name": "Ada", "lang": "Go", "expr": "a=b", "empty": ""}},
		{"a=1", map[string]string{"a": "1"}},
		{" =orphan; b = 2 ", map[string]string{"b": "2"}},
		{"", map[string]string{}},
	} {
		got := ParseKV(tt.in)
		if got == nil {
			t.Fatalf("ParseKV(%q) returned a nil map", tt.in)
		}
		if !maps.Equal(got, tt.want) {
			t.Errorf("ParseKV(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
