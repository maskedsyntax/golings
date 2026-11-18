// select_02: Until told to stop (reference solution)
package collect

// Collect receives from data until quit is closed and returns what it received.
func Collect(data <-chan int, quit <-chan struct{}) []int {
	var out []int
	for {
		select {
		case v := <-data:
			out = append(out, v)
		case <-quit:
			return out
		}
	}
}
