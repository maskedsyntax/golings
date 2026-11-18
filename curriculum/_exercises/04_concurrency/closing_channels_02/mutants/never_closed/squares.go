// Mutant: results is never closed.
package squares

import "sync"

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
	return results
}
