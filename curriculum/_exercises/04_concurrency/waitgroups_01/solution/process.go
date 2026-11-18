// waitgroups_01: Add before you go (reference solution)
package process

import "sync"

// ProcessAll runs work(i) for i in [0, n) concurrently and waits for all of them.
func ProcessAll(n int, work func(i int)) {
	var wg sync.WaitGroup
	for i := range n {
		wg.Add(1) // before the goroutine starts, so Wait can't miss it
		go func() {
			defer wg.Done()
			work(i)
		}()
	}
	wg.Wait()
}
