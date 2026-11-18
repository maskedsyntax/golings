// Mutant: adds one for the whole batch.
package process

import "sync"

func ProcessAll(n int, work func(i int)) {
	var wg sync.WaitGroup
	wg.Add(1)
	for i := range n {
		go func() {
			work(i)
			if i == 0 {
				wg.Done()
			}
		}()
	}
	wg.Wait()
}
