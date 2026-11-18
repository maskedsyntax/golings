// Mutant: Entries follows map iteration order.
package pairs

import "cmp"

type Pair[K, V any] struct {
	Key   K
	Value V
}

func Entries[K cmp.Ordered, V any](m map[K]V) []Pair[K, V] {
	var out []Pair[K, V]
	for k, v := range m {
		out = append(out, Pair[K, V]{k, v})
	}
	return out
}

func GroupBy[T any, K comparable](xs []T, key func(T) K) map[K][]T {
	groups := map[K][]T{}
	for _, x := range xs {
		k := key(x)
		groups[k] = append(groups[k], x)
	}
	return groups
}
