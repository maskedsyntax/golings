// closing_channels_01: No more values (reference solution)
package generate

// Generate sends each of nums on the returned channel, then closes it.
func Generate(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch) // the sender closes, after the last send
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}
