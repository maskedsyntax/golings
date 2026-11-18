// waitgroups_03: Collecting every error (reference solution)
package runall

import (
	"errors"
	"sync"
)

// RunAll runs every task concurrently and joins their errors, in task order.
func RunAll(tasks ...func() error) error {
	errs := make([]error, len(tasks))
	var wg sync.WaitGroup
	for i, task := range tasks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs[i] = task()
		}()
	}
	wg.Wait()
	return errors.Join(errs...)
}
