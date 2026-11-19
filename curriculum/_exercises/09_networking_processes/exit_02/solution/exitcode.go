// exit_02: Choosing exit codes (reference solution)
package exitcode

import (
	"context"
	"errors"
)

// UsageError reports a mistake in the command line.
type UsageError struct{ Msg string }

func (e *UsageError) Error() string { return "usage: " + e.Msg }

// ErrNotFound reports a missing resource.
var ErrNotFound = errors.New("not found")

// ExitCode maps an error returned by run to a process exit status.
func ExitCode(err error) int {
	var usage *UsageError
	switch {
	case err == nil:
		return 0
	case errors.As(err, &usage): // looks through wrapping
		return 2
	case errors.Is(err, ErrNotFound):
		return 3
	case errors.Is(err, context.Canceled):
		return 130
	default:
		return 1
	}
}
