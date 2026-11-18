// channels_02: Producer and consumer (reference solution)
package pipe

// Generate sends each of nums, in order, on the returned channel from a new goroutine.
func Generate(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

// Take receives exactly n values from ch.
func Take(ch <-chan int, n int) []int {
	out := make([]int, 0, n)
	for range n {
		out = append(out, <-ch)
	}
	return out
}
