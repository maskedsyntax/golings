// testing_and_benchmarking_01: Your first test
//
// Fix this test and add the missing case. See README.md.
//
// Reference: https://gobyexample.com/testing-and-benchmarking
package intmin

import "testing"

func TestIntMin(t testing.T) {
	if got := IntMin(2, -2); got != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2")
	}
	// TODO: also check a case where the first argument is the smaller one.
}
