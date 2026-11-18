// random_numbers_02: Random samples (reference solution)
package sample

import (
	"math/rand/v2"
	"slices"
)

// Sample returns k distinct items chosen uniformly at random using r.
func Sample(r *rand.Rand, items []string, k int) []string {
	k = max(0, min(k, len(items)))
	s := slices.Clone(items) // don't reorder the caller's slice
	r.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s[:k]
}
