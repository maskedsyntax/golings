// Mutant: only one sub-benchmark.
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
}
