// generics_04: Generic types and cmp.Ordered (reference solution)
package pairs

import (
	"cmp"
	"maps"
	"slices"
)

// Pair is a key and a value.
type Pair[K, V any] struct {
	Key   K
	Value V
}

// Entries returns the entries of m sorted by key.
func Entries[K cmp.Ordered, V any](m map[K]V) []Pair[K, V] {
	out := make([]Pair[K, V], 0, len(m))
	for _, k := range slices.Sorted(maps.Keys(m)) {
		out = append(out, Pair[K, V]{Key: k, Value: m[k]})
	}
	return out
}

// GroupBy groups xs by key(x), preserving order within each group.
func GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T {
	groups := map[K][]T{}
	for _, x := range xs {
		k := key(x)
		groups[k] = append(groups[k], x)
	}
	return groups
}
