// range_over_built_in_types_02: Ranging over ints and maps (reference solution)
package ranges

import "slices"

// Squares returns the squares of 1..n.
func Squares(n int) []int {
	var out []int
	for i := range n {
		out = append(out, (i+1)*(i+1))
	}
	return out
}

// SortedKeys returns the keys of m in ascending order.
func SortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
