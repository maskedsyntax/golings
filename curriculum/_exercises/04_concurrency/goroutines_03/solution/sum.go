// goroutines_03: A shared total (reference solution)
package sum

import "sync"

// ParallelSum returns the sum of every number in every chunk, summing each
// chunk in its own goroutine.
func ParallelSum(chunks [][]int) int {
	partial := make([]int, len(chunks)) // one slot per goroutine
	var wg sync.WaitGroup
	for i, chunk := range chunks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, n := range chunk {
				partial[i] += n
			}
		}()
	}
	wg.Wait()
	total := 0
	for _, p := range partial {
		total += p
	}
	return total
}
