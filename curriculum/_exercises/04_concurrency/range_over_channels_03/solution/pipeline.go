// range_over_channels_03: A pipeline that finishes (reference solution)
package pipeline

// Generate sends each number, then closes its output.
func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// Filter passes on the values for which keep returns true.
func Filter(in <-chan int, keep func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out) // in is drained: nothing more will be sent
		for v := range in {
			if keep(v) {
				out <- v
			}
		}
	}()
	return out
}

// Count returns how many values arrive before in is closed.
func Count(in <-chan int) int {
	n := 0
	for range in {
		n++
	}
	return n
}
