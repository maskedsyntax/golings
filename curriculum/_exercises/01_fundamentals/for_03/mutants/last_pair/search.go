// Mutant: break only leaves the inner loop, so later pairs overwrite earlier ones.
package search

func PositiveSum(xs []int) int {
	sum := 0
	for _, x := range xs {
		if x > 0 {
			sum += x
		}
	}
	return sum
}

func FirstPair(nums []int, target int) (i, j int, found bool) {
	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			if nums[a]+nums[b] == target {
				i, j, found = a, b, true
				break
			}
		}
	}
	return i, j, found
}
