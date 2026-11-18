// multiple_return_values_02: Minimum, maximum, and ok (reference solution)
package minmax

// MinMax returns the smallest and largest elements of nums.
// ok is false when nums is empty.
func MinMax(nums []int) (lo, hi int, ok bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}
	lo, hi = nums[0], nums[0]
	for _, n := range nums[1:] {
		lo = min(lo, n)
		hi = max(hi, n)
	}
	return lo, hi, true
}
