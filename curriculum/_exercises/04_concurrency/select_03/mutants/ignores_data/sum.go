// Mutant: returns before receiving anything.
package sum

func Sum(data <-chan int, done <-chan struct{}) int {
	<-done
	return 0
}
