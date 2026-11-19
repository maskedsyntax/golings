// Mutant: catches every bug but doesn't use t.Run.
package reverse

import "testing"

func TestReverse(t *testing.T) {
	for _, tt := range []struct{ in, want string }{{"abc", "cba"}, {"héllo", "olléh"}, {"", ""}} {
		if got := Reverse(tt.in); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
