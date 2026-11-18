// sorting_by_functions_01: Shortest first (reference solution)
package bylength

import (
	"cmp"
	"slices"
	"strings"
)

// ByLength sorts words in place, shortest first; words of equal length are
// in alphabetical order.
func ByLength(words []string) {
	slices.SortFunc(words, func(a, b string) int {
		return cmp.Or(
			cmp.Compare(len(a), len(b)), // by length...
			strings.Compare(a, b),       // ...then alphabetically
		)
	})
}
