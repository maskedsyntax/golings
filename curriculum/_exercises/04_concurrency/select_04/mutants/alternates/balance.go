// Mutant: alternates between a and b regardless of readiness.
package balance

func Balance(in <-chan int, a, b chan<- int, n int) {
	for i := range n {
		v := <-in
		if i%2 == 0 {
			a <- v
		} else {
			b <- v
		}
	}
}
