// Mutant: treating the zero value as "missing".
package maps

func WordCount(words []string) map[string]int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	return counts
}

func Has(m map[string]int, key string) bool {
	return m[key] != 0
}
