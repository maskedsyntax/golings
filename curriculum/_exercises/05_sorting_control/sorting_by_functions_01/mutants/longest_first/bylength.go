// Mutant: sorts longest first.
package bylength

import (
	"cmp"
	"slices"
	"strings"
)

func ByLength(words []string) {
	slices.SortFunc(words, func(a, b string) int {
		return cmp.Or(cmp.Compare(len(b), len(a)), strings.Compare(a, b))
	})
}
