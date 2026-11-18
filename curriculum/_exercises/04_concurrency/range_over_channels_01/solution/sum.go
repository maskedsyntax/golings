// range_over_channels_01: Ranging until closed (reference solution)
package sum

// Sum returns the total of every value received from ch until it is closed.
func Sum(ch <-chan int) int {
	total := 0
	for v := range ch {
		total += v
	}
	return total
}
