// errors_03: Checking wrapped errors (reference solution)
package config

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrPermission = errors.New("permission denied")
)

var store = map[string]string{"name": "golings", "locked": "secret"}

// Fetch returns the value stored under key.
func Fetch(key string) (string, error) {
	if key == "locked" {
		return "", fmt.Errorf("fetch %q: %w", key, ErrPermission)
	}
	v, ok := store[key]
	if !ok {
		return "", fmt.Errorf("fetch %q: %w", key, ErrNotFound)
	}
	return v, nil
}

// GetOrDefault returns the value for key, or def if the key does not exist.
// Other errors are returned to the caller.
func GetOrDefault(key, def string) (string, error) {
	v, err := Fetch(key)
	if errors.Is(err, ErrNotFound) {
		return def, nil
	}
	if err != nil {
		return "", err
	}
	return v, nil
}
