// errors_02: Sentinel errors and wrapping
//
// Implement Price. See README.md.
//
//	_, err := Price(prices, "kiwi")
//	errors.Is(err, ErrNotFound) == true
//	strings.Contains(err.Error(), "kiwi") == true
//
// Reference: https://gobyexample.com/errors
package price

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("invalid price")
)

// Price looks up key in prices.
func Price(prices map[string]int, key string) (int, error) {
	// TODO
	return 0, nil
}
