// pointers_03: Pointers to copies
//
// Writing through the returned pointers doesn't change nums. Fix it.
//
//	nums := []int{1, 2, 3}
//	*ElementPointers(nums)[1] = 20 // nums == [1 20 3]
//
// Reference: https://gobyexample.com/pointers
package elements

// ElementPointers returns a pointer to each element of nums, in order.
func ElementPointers(nums []int) []*int {
	var ptrs []*int
	for _, v := range nums {
		ptrs = append(ptrs, &v)
	}
	return ptrs
}
