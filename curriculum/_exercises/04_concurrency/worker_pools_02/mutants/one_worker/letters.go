// Mutant: avoids the deadlock by measuring everything in one goroutine.
package letters

func CountLetters(words []string, workers int, measure func(string) int) map[string]int {
	out := map[string]int{}
	for _, w := range words {
		out[w] = measure(w)
	}
	return out
}
