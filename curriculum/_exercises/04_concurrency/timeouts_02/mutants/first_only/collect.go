// Mutant: returns after the first value.
package collect

import "time"

func CollectFor(ch <-chan int, d time.Duration) []int {
	select {
	case v := <-ch:
		return []int{v}
	case <-time.After(d):
		return nil
	}
}
