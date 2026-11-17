// Mutant: reversing in place mutates the caller's slice.
package slices

func Reverse(in []int) []int {
	if in == nil {
		return []int{}
	}
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}
