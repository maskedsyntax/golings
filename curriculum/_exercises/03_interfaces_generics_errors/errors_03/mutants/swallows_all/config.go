// Mutant: any error at all turns into the default.
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

func GetOrDefault(key, def string) (string, error) {
	v, err := Fetch(key)
	if err != nil {
		return def, nil
	}
	return v, nil
}
