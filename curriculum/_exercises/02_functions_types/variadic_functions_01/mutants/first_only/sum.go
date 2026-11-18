// Mutant: Average only sums the first element.
package sum

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Average(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	return Sum(xs[0]) / len(xs)
}

func Total(a, b, c int) int { return Sum(a, b, c) }
