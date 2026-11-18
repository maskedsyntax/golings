// range_over_built_in_types_01: What range gives you (reference solution)
package ranges

// Sum returns the sum of the elements of nums.
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// CountRunes counts how often each character occurs in s.
func CountRunes(s string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}
	return counts
}
