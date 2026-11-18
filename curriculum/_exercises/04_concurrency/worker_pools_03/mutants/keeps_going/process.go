// Mutant: remembers the first error but processes every item anyway.
package process

import "sync"

func ProcessAll(items []int, workers int, f func(int) error) error {
	jobs := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer close(jobs)
		for _, it := range items {
			jobs <- it
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
					case errCh <- err:
					default:
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
