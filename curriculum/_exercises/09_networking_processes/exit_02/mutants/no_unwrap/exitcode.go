// Mutant: type assertion and ==, which miss wrapped errors.
package exitcode

import (
	"context"
	"errors"
)

type UsageError struct{ Msg string }

func (e *UsageError) Error() string { return "usage: " + e.Msg }

var ErrNotFound = errors.New("not found")

func ExitCode(err error) int {
	if err == nil {
		return 0
	}
	if _, ok := err.(*UsageError); ok {
		return 2
	}
	if err == ErrNotFound {
		return 3
	}
	if err == context.Canceled {
		return 130
	}
	return 1
}
