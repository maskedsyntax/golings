// testing_and_benchmarking_02: Table-driven tests (reference solution)
package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"ascii", "abc", "cba"},
		{"multi-byte", "héllo, 世界", "界世 ,olléh"},
		{"two", "ab", "ba"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.in); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
