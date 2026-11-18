package age

import (
	"strings"
	"testing"
)

func TestParseAgeValid(t *testing.T) {
	for _, tt := range []struct {
		s    string
		want int
	}{{"42", 42}, {" 7\n", 7}, {"0", 0}, {"150", 150}} {
		got, err := ParseAge(tt.s)
		if err != nil || got != tt.want {
			t.Errorf("ParseAge(%q) = %d, %v, want %d, nil", tt.s, got, err, tt.want)
		}
	}
}

func TestParseAgeInvalid(t *testing.T) {
	for _, tt := range []struct{ s, mention string }{
		{"abc", "abc"}, {"", ""}, {"4.5", "4.5"}, {"-1", "-1"}, {"151", "151"}, {"999", "999"},
	} {
		got, err := ParseAge(tt.s)
		if err == nil {
			t.Errorf("ParseAge(%q) = %d, nil, want an error", tt.s, got)
			continue
		}
		if got != 0 {
			t.Errorf("ParseAge(%q) returned %d alongside an error, want 0", tt.s, got)
		}
		if !strings.Contains(err.Error(), tt.mention) {
			t.Errorf("ParseAge(%q) error %q should mention %q", tt.s, err, tt.mention)
		}
	}
}
