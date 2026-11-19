// spawning_processes_01: Running a command
//
// Implement Run. See README.md.
//
// Reference: https://gobyexample.com/spawning-processes
package run

import (
	"context"
	"errors"
)

// Run executes name with args and returns its stdout and exit status.
func Run(ctx context.Context, name string, args ...string) (stdout string, code int, err error) {
	// TODO
	return "", -1, errors.New("not implemented")
}
