// random_numbers_01: Rolling dice
//
// Roll returns 0..5 instead of 1..6. Fix it. See README.md.
//
// Reference: https://gobyexample.com/random-numbers
package dice

import "math/rand/v2"

// Roll returns n rolls of a six-sided die, each between 1 and 6.
func Roll(r *rand.Rand, n int) []int {
	rolls := make([]int, n)
	for i := range rolls {
		rolls[i] = r.IntN(6)
	}
	return rolls
}
