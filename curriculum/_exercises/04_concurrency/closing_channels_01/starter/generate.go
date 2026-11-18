// closing_channels_01: No more values
//
// Receivers never learn that Generate is finished. Fix it. See README.md.
//
// Reference: https://gobyexample.com/closing-channels
package generate

// Generate sends each of nums on the returned channel, then closes it.
func Generate(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}
