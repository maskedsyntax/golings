// timeouts_02: An overall deadline
//
// CollectFor never stops while values keep arriving. Fix it. See README.md.
//
// Reference: https://gobyexample.com/timeouts
package collect

import "time"

// CollectFor returns the values received from ch during the next d.
func CollectFor(ch <-chan int, d time.Duration) []int {
	var out []int
	for {
		select {
		case v := <-ch:
			out = append(out, v)
		case <-time.After(d):
			return out
		}
	}
}
