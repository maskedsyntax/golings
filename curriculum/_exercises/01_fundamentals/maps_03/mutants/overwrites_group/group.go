// Mutant: assigning instead of appending keeps only the last word per length.
package group

import "unicode/utf8"

func GroupByLength(words []string) map[int][]string {
	groups := make(map[int][]string)
	for _, w := range words {
		groups[utf8.RuneCountInString(w)] = []string{w}
	}
	return groups
}
