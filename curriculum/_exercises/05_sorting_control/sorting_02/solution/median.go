// sorting_02: The median (reference solution)
package median

import "slices"

// Median returns the median of nums, or false if nums is empty.
// It does not modify nums.
func Median(nums []float64) (float64, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	s := slices.Clone(nums)
	slices.Sort(s)
	n := len(s)
	if n%2 == 1 {
		return s[n/2], true
	}
	return (s[n/2-1] + s[n/2]) / 2, true
}
