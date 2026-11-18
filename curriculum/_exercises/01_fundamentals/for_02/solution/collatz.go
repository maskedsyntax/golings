// for_02: Collatz steps (reference solution)
package collatz

// Steps returns the number of Collatz transformations needed to reach 1,
// or -1 for n <= 0.
func Steps(n int) int {
	if n <= 0 {
		return -1
	}
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps
}
