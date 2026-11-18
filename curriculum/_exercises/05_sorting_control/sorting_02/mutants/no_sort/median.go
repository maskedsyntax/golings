// Mutant: takes the middle of the unsorted input.
package median

func Median(nums []float64) (float64, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	n := len(nums)
	if n%2 == 1 {
		return nums[n/2], true
	}
	return (nums[n/2-1] + nums[n/2]) / 2, true
}
