// Mutant: Map overwrites the caller's slice; Filter reuses its backing array.
package apply

func Map(nums []int, f func(int) int) []int {
	for i, n := range nums {
		nums[i] = f(n)
	}
	return nums
}

func Filter(nums []int, keep func(int) bool) []int {
	out := nums[:0]
	for _, n := range nums {
		if keep(n) {
			out = append(out, n)
		}
	}
	return out
}
