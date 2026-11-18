// variadic_functions_02: At least one argument (reference solution)
package variadic

import "strings"

// Max returns the largest of its arguments.
func Max(first int, rest ...int) int {
	best := first
	for _, n := range rest {
		best = max(best, n)
	}
	return best
}

// JoinNonEmpty joins the non-empty parts with sep between them.
func JoinNonEmpty(sep string, parts ...string) string {
	var kept []string
	for _, p := range parts {
		if p != "" {
			kept = append(kept, p)
		}
	}
	return strings.Join(kept, sep)
}
