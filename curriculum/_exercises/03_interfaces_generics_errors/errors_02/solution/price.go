// errors_02: Sentinel errors and wrapping (reference solution)
package price

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("invalid price")
)

// Price looks up key in prices.
func Price(prices map[string]int, key string) (int, error) {
	v, ok := prices[key]
	if !ok {
		return 0, fmt.Errorf("price of %q: %w", key, ErrNotFound)
	}
	if v < 0 {
		return 0, fmt.Errorf("price of %q is %d: %w", key, v, ErrInvalid)
	}
	return v, nil
}
