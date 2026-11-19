// Mutant: returns the winner but leaves the others running.
package first

import (
	"context"
	"errors"
)

func FirstResult(ctx context.Context, fns ...func(context.Context) (string, error)) (string, error) {
	if len(fns) == 0 {
		return "", errors.New("no functions")
	}
	type result struct {
		s   string
		err error
	}
	results := make(chan result, len(fns))
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
