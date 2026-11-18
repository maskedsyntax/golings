// closing_channels_03: Stopping every worker
//
// stop hangs. Make it stop every worker and wait for them. See README.md.
//
// Reference: https://gobyexample.com/closing-channels
package workers

import "sync"

// Start runs n workers, each calling work(id) in a loop, until stop is called.
// stop returns once every worker has exited.
func Start(n int, work func(id int)) (stop func()) {
	quit := make(chan struct{})
	var wg sync.WaitGroup
	for id := range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-quit:
					return
				default:
					work(id)
				}
			}
		}()
	}
	return func() {
		quit <- struct{}{}
		wg.Wait()
	}
}
