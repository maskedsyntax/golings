// Mutant: subtests, but no empty-string case.
package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct{ name, in, want string }{
		{"ascii", "abc", "cba"},
		{"multi-byte", "héllo", "olléh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.in); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
