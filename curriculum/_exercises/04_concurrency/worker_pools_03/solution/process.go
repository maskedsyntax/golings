// worker_pools_03: Stopping on the first error (reference solution)
package process

import "sync"

// ProcessAll calls f on every item using the given number of workers. It
// returns the first error and stops dispatching new items once one occurs.
func ProcessAll(items []int, workers int, f func(int) error) error {
	jobs := make(chan int)
	quit := make(chan struct{})
	errCh := make(chan error, 1) // holds the first error only

	go func() {
		defer close(jobs)
		for _, it := range items {
			select {
			case jobs <- it:
			case <-quit:
				return
			}
		}
	}()

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for it := range jobs {
				if err := f(it); err != nil {
					select {
					case errCh <- err: // the first error: tell the dispatcher
						close(quit)
					default: // an earlier error already won
					}
				}
			}
		}()
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}
