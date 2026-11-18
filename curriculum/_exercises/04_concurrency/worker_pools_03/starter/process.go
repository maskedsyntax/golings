// worker_pools_03: Stopping on the first error
//
// Implement ProcessAll. See README.md.
//
// Reference: https://gobyexample.com/worker-pools
package process

// ProcessAll calls f on every item using the given number of workers. It
// returns the first error and stops dispatching new items once one occurs.
func ProcessAll(items []int, workers int, f func(int) error) error {
	// TODO
	return nil
}
