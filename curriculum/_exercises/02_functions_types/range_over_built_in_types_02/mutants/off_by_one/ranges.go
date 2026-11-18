// Mutant: squares 0..n-1 instead of 1..n.
package ranges

import "slices"

func Squares(n int) []int {
	var out []int
	for i := range n {
		out = append(out, i*i)
	}
	return out
}

func SortedKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
