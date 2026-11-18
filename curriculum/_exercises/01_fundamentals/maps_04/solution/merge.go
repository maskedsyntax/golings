// maps_04: Writing to a nil map (reference solution)
package merge

// Merge returns a new map containing the sum of a and b per key.
func Merge(a, b map[string]int) map[string]int {
	out := make(map[string]int, len(a)+len(b))
	for k, v := range a {
		out[k] += v
	}
	for k, v := range b {
		out[k] += v
	}
	return out
}
