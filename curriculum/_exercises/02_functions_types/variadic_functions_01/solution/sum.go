// variadic_functions_01: Spreading a slice (reference solution)
package sum

// Sum adds up its arguments.
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Average returns the integer mean of xs, or 0 if xs is empty.
func Average(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	return Sum(xs...) / len(xs)
}

// Total returns a + b + c, computed with Sum.
func Total(a, b, c int) int {
	return Sum(a, b, c)
}
