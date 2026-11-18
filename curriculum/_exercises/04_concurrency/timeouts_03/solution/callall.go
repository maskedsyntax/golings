// timeouts_03: A timeout for each call (reference solution)
package callall

import (
	"sync"
	"time"
)

// Result is the outcome of one call.
type Result struct {
	Value    int
	TimedOut bool
}

// CallAll runs every call, allowing each at most d, and returns their results in order.
func CallAll(calls []func() int, d time.Duration) []Result {
	results := make([]Result, len(calls))
	var wg sync.WaitGroup
	for i, call := range calls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch := make(chan int, 1)
			go func() { ch <- call() }()
			select {
			case v := <-ch:
				results[i] = Result{Value: v}
			case <-time.After(d):
				results[i] = Result{TimedOut: true}
			}
		}()
	}
	wg.Wait()
	return results
}
