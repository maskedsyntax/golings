// timeouts_02: An overall deadline (reference solution)
package collect

import "time"

// CollectFor returns the values received from ch during the next d.
func CollectFor(ch <-chan int, d time.Duration) []int {
	var out []int
	deadline := time.After(d) // one timer for the whole call
	for {
		select {
		case v := <-ch:
			out = append(out, v)
		case <-deadline:
			return out
		}
	}
}
