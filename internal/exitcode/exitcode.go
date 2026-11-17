// Package exitcode defines the process exit codes used by the golings CLI
// (spec §3.3) and an error type that carries one.
package exitcode

import (
	"errors"
	"fmt"
)

const (
	OK           = 0   // success
	VerifyFailed = 1   // exercise verification failure
	Usage        = 2   // command/argument error
	Workspace    = 3   // invalid workspace or configuration
	Toolchain    = 4   // missing/incompatible Go toolchain
	Infra        = 5   // infrastructure or time-limit failure
	Interrupted  = 130 // interrupted by Ctrl+C
)

// Error is an error annotated with the exit code the CLI should use.
type Error struct {
	Code int
	Err  error
}

func (e *Error) Error() string { return e.Err.Error() }
func (e *Error) Unwrap() error { return e.Err }

// New returns an *Error with a formatted message.
func New(code int, format string, args ...any) error {
	return &Error{Code: code, Err: fmt.Errorf(format, args...)}
}

// Wrap annotates err with code. It returns nil when err is nil.
func Wrap(code int, err error) error {
	if err == nil {
		return nil
	}
	return &Error{Code: code, Err: err}
}

// Code extracts the exit code for err. Errors without an annotation map to
// Infra, because they indicate something unexpected went wrong.
func Code(err error) int {
	if err == nil {
		return OK
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return Infra
}
