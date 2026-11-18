// Mutant: reusing a as the result changes the caller's map.
package merge

func Merge(a, b map[string]int) map[string]int {
	out := a
	if out == nil {
		out = map[string]int{}
	}
	for k, v := range b {
		out[k] += v
	}
	return out
}
