// Mutant: returns keys in map iteration order.
package ranges

func Squares(n int) []int {
	var out []int
	for i := range n {
		out = append(out, (i+1)*(i+1))
	}
	return out
}

func SortedKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
