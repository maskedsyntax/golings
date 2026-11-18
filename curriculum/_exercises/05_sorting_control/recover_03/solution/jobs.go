// recover_03: A panic in a goroutine (reference solution)
package jobs

import (
	"fmt"
	"sync"
)

// RunJobs runs every job concurrently. errs[i] reports a panic in jobs[i].
func RunJobs(jobs []func()) (errs []error) {
	errs = make([]error, len(jobs))
	var wg sync.WaitGroup
	for i, job := range jobs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// recover only works in the goroutine that panicked.
			defer func() {
				if r := recover(); r != nil {
					errs[i] = fmt.Errorf("job %d panicked: %v", i, r)
				}
			}()
			job()
		}()
	}
	wg.Wait()
	return errs
}
