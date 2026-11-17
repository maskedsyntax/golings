// maps_01: Counting and membership (reference solution)
package maps

// WordCount returns how many times each word occurs (case-sensitive).
// It never returns a nil map.
func WordCount(words []string) map[string]int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	return counts
}

// Has reports whether key is present in m, even when its value is zero.
func Has(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}
