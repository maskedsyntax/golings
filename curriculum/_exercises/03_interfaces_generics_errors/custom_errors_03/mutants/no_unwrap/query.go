// Mutant: no Unwrap method, so the cause is hidden from errors.Is.
package query

import (
	"errors"
	"fmt"
)

var (
	ErrTimeout = errors.New("timeout")
	ErrSyntax  = errors.New("syntax error")
)

func exec(q string) error {
	switch q {
	case "slow":
		return ErrTimeout
	case "":
		return fmt.Errorf("empty query: %w", ErrSyntax)
	}
	return nil
}

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string { return fmt.Sprintf("query %q: %v", e.Query, e.Err) }

func Run(q string) error {
	if err := exec(q); err != nil {
		return &QueryError{Query: q, Err: err}
	}
	return nil
}
