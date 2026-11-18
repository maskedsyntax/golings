// Mutant: 1..5.
package dice

import "math/rand/v2"

func Roll(r *rand.Rand, n int) []int {
	rolls := make([]int, n)
	for i := range rolls {
		rolls[i] = r.IntN(5) + 1
	}
	return rolls
}
