// context_02: First answer wins (reference solution)
package first

import (
	"context"
	"errors"
)

// FirstResult runs fns concurrently and returns the first success, cancelling the rest.
func FirstResult(ctx context.Context, fns ...func(context.Context) (string, error)) (string, error) {
	if len(fns) == 0 {
		return "", errors.New("FirstResult: no functions")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // stops the losers once we return

	type result struct {
		s   string
		err error
	}
	results := make(chan result, len(fns)) // buffered: late senders never block
	for _, fn := range fns {
		go func() {
			s, err := fn(ctx)
			results <- result{s, err}
		}()
	}

	var errs []error
	for range fns {
		r := <-results
		if r.err == nil {
			return r.s, nil
		}
		errs = append(errs, r.err)
	}
	if err := ctx.Err(); err != nil {
		return "", err
	}
	return "", errors.Join(errs...)
}
