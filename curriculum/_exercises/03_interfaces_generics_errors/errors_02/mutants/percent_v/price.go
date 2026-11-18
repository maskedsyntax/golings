// Mutant: %v copies the message but doesn't wrap the sentinel.
package price

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("invalid price")
)

func Price(prices map[string]int, key string) (int, error) {
	v, ok := prices[key]
	if !ok {
		return 0, fmt.Errorf("price of %q: %v", key, ErrNotFound)
	}
	if v < 0 {
		return 0, fmt.Errorf("price of %q is %d: %v", key, v, ErrInvalid)
	}
	return v, nil
}
