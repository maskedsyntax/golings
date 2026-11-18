// worker_pools_04: Fan-in
//
// Implement Merge. See README.md.
//
// Reference: https://gobyexample.com/worker-pools
package merge

// Merge copies every value from every input to the returned channel and
// closes it once all inputs are closed.
func Merge(chans ...<-chan int) <-chan int {
	// TODO
	return nil
}
