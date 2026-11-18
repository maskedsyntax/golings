// functions_02: Functions as values (reference solution)
package apply

// Map returns a new slice holding f(n) for every n in nums.
func Map(nums []int, f func(int) int) []int {
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		out = append(out, f(n))
	}
	return out
}

// Filter returns a new slice holding the elements of nums for which keep is true.
func Filter(nums []int, keep func(int) bool) []int {
	var out []int
	for _, n := range nums {
		if keep(n) {
			out = append(out, n)
		}
	}
	return out
}
