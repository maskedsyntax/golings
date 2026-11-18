// closing_channels_02: Many senders, one close
//
// SquareAll panics. Close results exactly once. See README.md.
//
// Reference: https://gobyexample.com/closing-channels
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
			defer close(results)
			for n := range jobs {
				results <- n * n
			}
		}()
	}
	return results
}
