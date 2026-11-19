// context_02: First answer wins
//
// Implement FirstResult. See README.md.
//
// Reference: https://gobyexample.com/context
package first

import (
	"context"
	"errors"
)

// FirstResult runs fns concurrently and returns the first success, cancelling the rest.
func FirstResult(ctx context.Context, fns ...func(context.Context) (string, error)) (string, error) {
	// TODO
	return "", errors.New("not implemented")
}
