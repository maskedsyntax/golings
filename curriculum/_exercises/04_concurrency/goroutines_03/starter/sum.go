// goroutines_03: A shared total
//
// ParallelSum has a data race. Remove it without a mutex. See README.md.
//
// Reference: https://gobyexample.com/goroutines
package sum

import "sync"

// ParallelSum returns the sum of every number in every chunk, summing each
// chunk in its own goroutine.
func ParallelSum(chunks [][]int) int {
	total := 0
	var wg sync.WaitGroup
	for _, chunk := range chunks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, n := range chunk {
				total += n
			}
		}()
	}
	wg.Wait()
	return total
}
