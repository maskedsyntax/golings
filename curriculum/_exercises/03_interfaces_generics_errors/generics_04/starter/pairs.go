// generics_04: Generic types and cmp.Ordered
//
// Implement Entries and GroupBy. See README.md.
//
//	Entries(map[string]int{"b": 2, "a": 1}) == []Pair[string, int]{{"a", 1}, {"b", 2}}
//
// Reference: https://gobyexample.com/generics
package pairs

import "cmp"

// Pair is a key and a value.
type Pair[K, V any] struct {
	Key   K
	Value V
}

// Entries returns the entries of m sorted by key.
func Entries[K cmp.Ordered, V any](m map[K]V) []Pair[K, V] {
	// TODO
	return nil
}

// GroupBy groups xs by key(x), preserving order within each group.
func GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T {
	// TODO
	return nil
}
