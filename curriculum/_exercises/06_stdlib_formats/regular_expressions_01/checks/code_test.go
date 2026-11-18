package code

import "testing"

func TestIsValidCode(t *testing.T) {
	for _, tt := range []struct {
		s    string
		want bool
	}{
		{"AB-1234", true},
		{"ZZ-0000", true},
		{"ab-1234", false},
		{"AB-123", false},
		{"xAB-1234", false},
		{"AB-12345", false},
		{"AB-1234\n", false},
		{"", false},
	} {
		if got := IsValidCode(tt.s); got != tt.want {
			t.Errorf("IsValidCode(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
