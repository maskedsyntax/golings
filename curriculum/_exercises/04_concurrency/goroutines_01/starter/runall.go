// goroutines_01: Starting goroutines
//
// RunAll should run every task concurrently and return when all are done.
// Right now it calls them one at a time. See README.md.
//
// Reference: https://gobyexample.com/goroutines
package runall

import "sync"

// RunAll runs every task in its own goroutine and waits for all of them.
func RunAll(tasks ...func()) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		// TODO: run this in a new goroutine.
		task()
		wg.Done()
	}
	wg.Wait()
}
