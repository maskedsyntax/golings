package truncate

import (
	"testing"
	"unicode/utf8"
)

func TestTruncate(t *testing.T) {
	for _, tt := range []struct {
		s    string
		max  int
		want string
	}{
		{"hello world", 6, "hello…"},
		{"hello", 5, "hello"},
		{"hello", 10, "hello"},
		{"crème brûlée", 4, "crè…"},
		{"日本語のテキスト", 3, "日本…"},
		{"abc", 1, "…"},
		{"abc", 0, ""},
		{"abc", -2, ""},
		{"", 3, ""},
	} {
		got := Truncate(tt.s, tt.max)
		if !utf8.ValidString(got) {
			t.Errorf("Truncate(%q, %d) = %q, which is not valid UTF-8", tt.s, tt.max, got)
		} else if got != tt.want {
			t.Errorf("Truncate(%q, %d) = %q, want %q", tt.s, tt.max, got, tt.want)
		}
	}
}
