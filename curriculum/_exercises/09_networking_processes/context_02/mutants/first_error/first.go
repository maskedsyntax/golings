// Mutant: returns the first result even if it is an error.
package first

import (
	"context"
	"errors"
)

func FirstResult(ctx context.Context, fns ...func(context.Context) (string, error)) (string, error) {
	if len(fns) == 0 {
		return "", errors.New("no functions")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
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
	r := <-results
	return r.s, r.err
}
