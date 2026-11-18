// slices_03: Remove duplicate values (reference solution)
package unique

import "slices"

// Unique returns the distinct values of numbers in first-seen order.
func Unique(numbers []int) []int {
	out := []int{}
	for _, n := range numbers {
		if !slices.Contains(out, n) {
			out = append(out, n)
		}
	}
	return out
}
