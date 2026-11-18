// Mutant: recovers in each goroutine but always reports in errs[0].
package jobs

import (
	"fmt"
	"sync"
)

func RunJobs(jobs []func()) (errs []error) {
	errs = make([]error, len(jobs))
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					mu.Lock()
					errs[0] = fmt.Errorf("a job panicked: %v", r)
					mu.Unlock()
				}
			}()
			job()
		}()
	}
	wg.Wait()
	return errs
}
