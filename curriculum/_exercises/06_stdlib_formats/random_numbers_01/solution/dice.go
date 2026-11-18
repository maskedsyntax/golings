// random_numbers_01: Rolling dice (reference solution)
package dice

import "math/rand/v2"

// Roll returns n rolls of a six-sided die, each between 1 and 6.
func Roll(r *rand.Rand, n int) []int {
	rolls := make([]int, n)
	for i := range rolls {
		rolls[i] = r.IntN(6) + 1 // IntN(6) is 0..5
	}
	return rolls
}
