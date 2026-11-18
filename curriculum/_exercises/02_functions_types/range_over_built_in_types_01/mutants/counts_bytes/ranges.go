// Mutant: Sum is fixed, but CountRunes still counts bytes.
package ranges

func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func CountRunes(s string) map[rune]int {
	counts := map[rune]int{}
	for i := 0; i < len(s); i++ {
		counts[rune(s[i])]++
	}
	return counts
}
