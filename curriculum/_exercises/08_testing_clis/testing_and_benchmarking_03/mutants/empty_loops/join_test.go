// Mutant: right names, but the loops call nothing.
package join

import "testing"

func BenchmarkJoin(b *testing.B) {
	b.Run("plus", func(b *testing.B) {
		for b.Loop() {
		}
	})
	b.Run("builder", func(b *testing.B) {
		for b.Loop() {
		}
	})
}
