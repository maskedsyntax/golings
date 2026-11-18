// Mutant: returns the sentinels without context.
package price

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("invalid price")
)

func Price(prices map[string]int, key string) (int, error) {
	v, ok := prices[key]
	if !ok {
		return 0, ErrNotFound
	}
	if v < 0 {
		return 0, ErrInvalid
	}
	return v, nil
}
