// channel_directions_02: A pipeline stage (reference solution)
package square

// Square returns a channel that delivers the square of every value received from in.
func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			v := <-in
			out <- v * v
		}
	}()
	return out // narrowed to <-chan int by the return type
}
