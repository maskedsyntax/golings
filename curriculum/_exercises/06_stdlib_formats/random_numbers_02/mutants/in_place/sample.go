// Mutant: shuffles the caller's slice.
package sample

import "math/rand/v2"

func Sample(r *rand.Rand, items []string, k int) []string {
	k = max(0, min(k, len(items)))
	r.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	return items[:k]
}
