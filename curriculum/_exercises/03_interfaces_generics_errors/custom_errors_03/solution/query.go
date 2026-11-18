// custom_errors_03: Unwrap (reference solution)
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

func (e *QueryError) Error() string {
	return fmt.Sprintf("query %q: %v", e.Query, e.Err)
}

// Unwrap lets errors.Is and errors.As see the cause.
func (e *QueryError) Unwrap() error {
	return e.Err
}

// Run executes q and wraps any failure in a *QueryError.
func Run(q string) error {
	if err := exec(q); err != nil {
		return &QueryError{Query: q, Err: err}
	}
	return nil
}
