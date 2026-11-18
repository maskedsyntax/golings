// Mutant: points into a copy of nums.
package elements

import "slices"

func ElementPointers(nums []int) []*int {
	cp := slices.Clone(nums)
	var ptrs []*int
	for i := range cp {
		ptrs = append(ptrs, &cp[i])
	}
	return ptrs
}
