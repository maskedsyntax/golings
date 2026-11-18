// select_04: Sending to whoever is free (reference solution)
package balance

// Balance moves n values from in to whichever of a or b is ready to receive.
func Balance(in <-chan int, a, b chan<- int, n int) {
	for range n {
		v := <-in
		select {
		case a <- v:
		case b <- v:
		}
	}
}
