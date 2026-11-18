// waitgroups_03: Collecting every error
//
// RunAll runs tasks one at a time. Make it concurrent. See README.md.
//
// Reference: https://gobyexample.com/waitgroups
package runall

import "errors"

// RunAll runs every task concurrently and joins their errors, in task order.
func RunAll(tasks ...func() error) error {
	var errs []error
	for _, task := range tasks {
		errs = append(errs, task())
	}
	return errors.Join(errs...)
}
