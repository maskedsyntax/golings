// Mutant: Max starts at 0; JoinNonEmpty keeps empty parts.
package variadic

import "strings"

func Max(first int, rest ...int) int {
	best := 0
	for _, n := range append([]int{first}, rest...) {
		best = max(best, n)
	}
	return best
}

func JoinNonEmpty(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}
