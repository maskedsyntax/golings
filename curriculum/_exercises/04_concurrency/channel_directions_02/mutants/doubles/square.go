// Mutant: doubles instead of squaring.
package square

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			out <- 2 * <-in
		}
	}()
	return out
}
