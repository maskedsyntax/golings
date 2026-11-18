// channels_02: Producer and consumer
//
// Implement Generate and Take. See README.md.
//
// Reference: https://gobyexample.com/channels
package pipe

// Generate sends each of nums, in order, on the returned channel from a new goroutine.
func Generate(nums ...int) <-chan int {
	// TODO
	return nil
}

// Take receives exactly n values from ch.
func Take(ch <-chan int, n int) []int {
	// TODO
	return nil
}
