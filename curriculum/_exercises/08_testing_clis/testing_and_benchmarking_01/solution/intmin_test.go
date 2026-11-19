// testing_and_benchmarking_01: Your first test (reference solution)
package intmin

import "testing"

func TestIntMin(t *testing.T) {
	if got := IntMin(2, -2); got != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", got)
	}
	if got := IntMin(-3, 5); got != -3 {
		t.Errorf("IntMin(-3, 5) = %d; want -3", got)
	}
}
