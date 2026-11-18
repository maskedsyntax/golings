// Mutant: not random at all.
package sample

import (
	"math/rand/v2"
	"slices"
)

func Sample(r *rand.Rand, items []string, k int) []string {
	k = max(0, min(k, len(items)))
	return slices.Clone(items[:k])
}
