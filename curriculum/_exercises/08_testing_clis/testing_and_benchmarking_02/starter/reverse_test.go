// testing_and_benchmarking_02: Table-driven tests
//
// Finish the table-driven test. See README.md.
//
// Reference: https://gobyexample.com/testing-and-benchmarking
package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"ascii", "abc", "cba"},
		// TODO: add more cases.
	}
	for _, tt := range tests {
		// TODO: run each case as a subtest with t.Run(tt.name, ...).
		if got := Reverse(tt.in); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
