// custom_errors_03: Unwrap
//
// Implement QueryError's methods and Run. See README.md.
//
//	err := Run("slow")
//	errors.Is(err, ErrTimeout) == true
//	err.Error() == `query "slow": timeout`
//
// Reference: https://gobyexample.com/custom-errors
package query

import (
	"errors"
	"fmt"
)

var (
	ErrTimeout = errors.New("timeout")
	ErrSyntax  = errors.New("syntax error")
)

// exec pretends to run a query.
func exec(q string) error {
	switch q {
	case "slow":
		return ErrTimeout
	case "":
		return fmt.Errorf("empty query: %w", ErrSyntax)
	}
	return nil
}

// QueryError records which query failed and why.
type QueryError struct {
	Query string
	Err   error
}

// TODO: func (e *QueryError) Error() string
// TODO: func (e *QueryError) Unwrap() error

// Run executes q and wraps any failure in a *QueryError.
func Run(q string) error {
	// TODO
	return nil
}
