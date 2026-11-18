// select_03: Breaking out of select (reference solution)
package sum

// Sum adds every value received from data until done is closed.
func Sum(data <-chan int, done <-chan struct{}) int {
	total := 0
	for {
		select {
		case v := <-data:
			total += v
		case <-done:
			return total
		}
	}
}
