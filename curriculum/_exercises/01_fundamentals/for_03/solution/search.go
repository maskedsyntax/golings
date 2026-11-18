// for_03: break, continue and labels (reference solution)
package search

// PositiveSum returns the sum of the positive elements of xs.
func PositiveSum(xs []int) int {
	sum := 0
	for _, x := range xs {
		if x <= 0 {
			continue
		}
		sum += x
	}
	return sum
}

// FirstPair returns the first indices i < j whose values sum to target.
func FirstPair(nums []int, target int) (i, j int, found bool) {
outer:
	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			if nums[a]+nums[b] == target {
				i, j, found = a, b, true
				break outer
			}
		}
	}
	return i, j, found
}
