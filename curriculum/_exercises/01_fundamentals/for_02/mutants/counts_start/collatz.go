// Mutant: counts the numbers visited instead of the transformations.
package collatz

func Steps(n int) int {
	if n <= 0 {
		return -1
	}
	count := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		count++
	}
	return count
}
