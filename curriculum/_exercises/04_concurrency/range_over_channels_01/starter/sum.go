// range_over_channels_01: Ranging until closed
//
// Sum assumes exactly three values. Use for ... range. See README.md.
//
// Reference: https://gobyexample.com/range-over-channels
package sum

// Sum returns the total of every value received from ch until it is closed.
func Sum(ch <-chan int) int {
	total := 0
	for i := 0; i < 3; i++ {
		total += <-ch
	}
	return total
}
