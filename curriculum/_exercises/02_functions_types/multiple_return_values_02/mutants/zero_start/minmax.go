// Mutant: starts from zero, so all-positive or all-negative input is wrong.
package minmax

func MinMax(nums []int) (lo, hi int, ok bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}
	for _, n := range nums {
		lo = min(lo, n)
		hi = max(hi, n)
	}
	return lo, hi, true
}
