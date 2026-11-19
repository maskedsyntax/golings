// testing_and_benchmarking_03: Benchmarks
//
// Write the sub-benchmarks. See README.md.
//
// Reference: https://gobyexample.com/testing-and-benchmarking
package join

import (
	"strings"
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	parts := strings.Fields(strings.Repeat("gopher ", 200))
	// TODO: two sub-benchmarks, "plus" and "builder".
	_ = parts
}
