// Mutant: compiles, but only checks the case "return b" also passes.
package intmin

import "testing"

func TestIntMin(t *testing.T) {
	if got := IntMin(2, -2); got != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", got)
	}
}
