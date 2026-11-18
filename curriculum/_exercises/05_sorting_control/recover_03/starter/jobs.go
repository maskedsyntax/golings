// recover_03: A panic in a goroutine
//
// One panicking job crashes everything. Fix RunJobs. See README.md.
//
// Reference: https://gobyexample.com/recover
package jobs

import (
	"fmt"
	"sync"
)

// RunJobs runs every job concurrently. errs[i] reports a panic in jobs[i].
func RunJobs(jobs []func()) (errs []error) {
	errs = make([]error, len(jobs))
	defer func() {
		if r := recover(); r != nil {
			errs[0] = fmt.Errorf("a job panicked: %v", r)
		}
	}()
	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			job()
		}()
	}
	wg.Wait()
	return errs
}
