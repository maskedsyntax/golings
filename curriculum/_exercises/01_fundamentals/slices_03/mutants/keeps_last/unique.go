// Mutant: sorting and compacting loses first-seen order.
package unique

import "slices"

func Unique(numbers []int) []int {
	out := append([]int{}, numbers...)
	slices.Sort(out)
	return slices.Compact(out)
}
