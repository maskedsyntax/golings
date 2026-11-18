// maps_03: Grouping into a map of slices (reference solution)
package group

import "unicode/utf8"

// GroupByLength groups words by rune length, preserving input order.
func GroupByLength(words []string) map[int][]string {
	groups := make(map[int][]string)
	for _, w := range words {
		n := utf8.RuneCountInString(w)
		groups[n] = append(groups[n], w)
	}
	return groups
}
