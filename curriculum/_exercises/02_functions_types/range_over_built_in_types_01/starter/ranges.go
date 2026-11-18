// range_over_built_in_types_01: What range gives you
//
// Both functions return wrong answers. Fix them.
//
//	Sum([]int{5, 6, 7}) == 18
//	CountRunes("héé") == map[rune]int{'h': 1, 'é': 2}
//
// Reference: https://gobyexample.com/range-over-built-in-types
package ranges

// Sum returns the sum of the elements of nums.
func Sum(nums []int) int {
	total := 0
	for n := range nums {
		total += n
	}
	return total
}

// CountRunes counts how often each character occurs in s.
func CountRunes(s string) map[rune]int {
	counts := map[rune]int{}
	for i := range s {
		counts[rune(s[i])]++
	}
	return counts
}
