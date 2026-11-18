// worker_pools_01: A fixed number of workers (reference solution)
package pool

import "sync"

// MapPool returns f(items[i]) for every i, using at most workers goroutines.
func MapPool(items []int, workers int, f func(int) int) []int {
	results := make([]int, len(items))
	jobs := make(chan int)
	go func() {
		defer close(jobs)
		for i := range items {
			jobs <- i
		}
	}()
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				results[i] = f(items[i])
			}
		}()
	}
	wg.Wait()
	return results
}
