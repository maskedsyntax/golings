package query

import (
	"errors"
	"fmt"
	"testing"
)

func TestRunOK(t *testing.T) {
	if err := Run("select 1"); err != nil {
		t.Errorf("Run(\"select 1\") = %v, want nil", err)
	}
}

func TestRunTimeout(t *testing.T) {
	err := Run("slow")
	var qe *QueryError
	if !errors.As(err, &qe) {
		t.Fatalf("Run(\"slow\") = %v (%T), want a *QueryError", err, err)
	}
	if qe.Query != "slow" {
		t.Errorf("Query = %q, want \"slow\"", qe.Query)
	}
	if !errors.Is(err, ErrTimeout) {
		t.Error("errors.Is(err, ErrTimeout) = false; does QueryError have an Unwrap method?")
	}
	if got := err.Error(); got != `query "slow": timeout` {
		t.Errorf("Error() = %q, want %q", got, `query "slow": timeout`)
	}
}

func TestDeepChain(t *testing.T) {
	err := fmt.Errorf("handler: %w", Run(""))
	if !errors.Is(err, ErrSyntax) {
		t.Errorf("errors.Is(%v, ErrSyntax) = false; the chain is handler -> QueryError -> empty query -> ErrSyntax", err)
	}
	if errors.Is(err, ErrTimeout) {
		t.Error("a syntax error should not match ErrTimeout")
	}
	if got := err.Error(); got != `handler: query "": empty query: syntax error` {
		t.Errorf("Error() = %q", got)
	}
}
