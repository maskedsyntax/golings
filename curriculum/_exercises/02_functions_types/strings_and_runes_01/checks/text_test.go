package text

import (
	"testing"
	"unicode/utf8"
)

func TestLength(t *testing.T) {
	for _, tt := range []struct {
		s    string
		want int
	}{{"hello", 5}, {"héllo", 5}, {"日本", 2}, {"", 0}, {"🙂!", 2}} {
		if got := Length(tt.s); got != tt.want {
			t.Errorf("Length(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, tt := range []struct{ s, want string }{
		{"abc", "cba"}, {"añb", "bña"}, {"日本語", "語本日"}, {"", ""}, {"x🙂y", "y🙂x"},
	} {
		got := Reverse(tt.s)
		if !utf8.ValidString(got) {
			t.Errorf("Reverse(%q) = %q, which is not valid UTF-8", tt.s, got)
		} else if got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
		}
	}
}
