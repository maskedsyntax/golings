// slices_01: Reverse without mutating (reference solution)
package slices

// Reverse returns a reversed copy of in.
func Reverse(in []int) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[len(in)-1-i] = v
	}
	return out
}
