// timeouts_03: A timeout for each call
//
// CallAll is correct but slow: the timeouts add up. Make the calls
// concurrent. See README.md.
//
// Reference: https://gobyexample.com/timeouts
package callall

import "time"

// Result is the outcome of one call.
type Result struct {
	Value    int
	TimedOut bool
}

// CallAll runs every call, allowing each at most d, and returns their results in order.
func CallAll(calls []func() int, d time.Duration) []Result {
	results := make([]Result, len(calls))
	for i, call := range calls {
		ch := make(chan int, 1)
		go func() { ch <- call() }()
		select {
		case v := <-ch:
			results[i] = Result{Value: v}
		case <-time.After(d):
			results[i] = Result{TimedOut: true}
		}
	}
	return results
}
