// Mutant: one worker regardless of the requested number.
package pool

func MapPool(items []int, workers int, f func(int) int) []int {
	results := make([]int, len(items))
	for i, item := range items {
		results[i] = f(item)
	}
	return results
}
