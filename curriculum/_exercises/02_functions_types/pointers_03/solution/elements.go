// pointers_03: Pointers to copies (reference solution)
package elements

// ElementPointers returns a pointer to each element of nums, in order.
func ElementPointers(nums []int) []*int {
	ptrs := make([]*int, 0, len(nums))
	for i := range nums {
		ptrs = append(ptrs, &nums[i]) // the element's own storage
	}
	return ptrs
}
