// Mutant: one shared deadline, checked sequentially; later calls get no time.
package callall

import "time"

type Result struct {
	Value    int
	TimedOut bool
}

func CallAll(calls []func() int, d time.Duration) []Result {
	results := make([]Result, len(calls))
	deadline := time.After(d)
	expired := false
	for i, call := range calls {
		if expired {
			results[i] = Result{TimedOut: true}
			continue
		}
		ch := make(chan int, 1)
		go func() { ch <- call() }()
		select {
		case v := <-ch:
			results[i] = Result{Value: v}
		case <-deadline:
			results[i] = Result{TimedOut: true}
			expired = true
		}
	}
	return results
}
