// select_03: Breaking out of select
//
// Sum never returns. Fix it. See README.md.
//
// Reference: https://gobyexample.com/select
package sum

// Sum adds every value received from data until done is closed.
func Sum(data <-chan int, done <-chan struct{}) int {
	total := 0
	for {
		select {
		case v := <-data:
			total += v
		case <-done:
			break
		}
	}
	return total
}
