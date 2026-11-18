// waitgroups_01: Add before you go
//
// ProcessAll returns before its goroutines finish. Fix it. See README.md.
//
// Reference: https://gobyexample.com/waitgroups
package process

import "sync"

// ProcessAll runs work(i) for i in [0, n) concurrently and waits for all of them.
func ProcessAll(n int, work func(i int)) {
	var wg sync.WaitGroup
	for i := range n {
		go func() {
			wg.Add(1)
			defer wg.Done()
			work(i)
		}()
	}
	wg.Wait()
}
