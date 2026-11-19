// exit_02: Choosing exit codes
//
// Implement ExitCode. See README.md.
//
// Reference: https://gobyexample.com/exit
package exitcode

import "errors"

// UsageError reports a mistake in the command line.
type UsageError struct{ Msg string }

func (e *UsageError) Error() string { return "usage: " + e.Msg }

// ErrNotFound reports a missing resource.
var ErrNotFound = errors.New("not found")

// ExitCode maps an error returned by run to a process exit status.
func ExitCode(err error) int {
	// TODO
	return 0
}
