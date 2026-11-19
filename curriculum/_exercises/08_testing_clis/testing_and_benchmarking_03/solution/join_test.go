// testing_and_benchmarking_03: Benchmarks (reference solution)
package join

import (
	"strings"
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	parts := strings.Fields(strings.Repeat("gopher ", 200))
	b.Run("plus", func(b *testing.B) {
		for b.Loop() {
			JoinPlus(parts)
		}
	})
	b.Run("builder", func(b *testing.B) {
		for b.Loop() {
			JoinBuilder(parts)
		}
	})
}
