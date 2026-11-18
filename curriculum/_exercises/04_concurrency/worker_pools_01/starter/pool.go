// worker_pools_01: A fixed number of workers
//
// MapPool starts a goroutine per item. Use a pool of workers. See README.md.
//
// Reference: https://gobyexample.com/worker-pools
package pool

import "sync"

// MapPool returns f(items[i]) for every i, using at most workers goroutines.
func MapPool(items []int, workers int, f func(int) int) []int {
	results := make([]int, len(items))
	var wg sync.WaitGroup
	for i, item := range items {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results[i] = f(item)
		}()
	}
	wg.Wait()
	return results
}
