// closing_channels_02: Many senders, one close (reference solution)
package squares

import "sync"

// SquareAll returns a channel delivering n*n for every n (in any order),
// computed by the given number of workers, and closed when all are sent.
func SquareAll(nums []int, workers int) <-chan int {
	jobs := make(chan int)
	results := make(chan int)
	go func() {
		defer close(jobs)
		for _, n := range nums {
			jobs <- n
		}
	}()
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				results <- n * n
			}
		}()
	}
	go func() {
		wg.Wait() // every worker has sent its last result
		close(results)
	}()
	return results
}
