// sorting_by_functions_01: Shortest first
//
// ByLength sorts alphabetically. Sort by length instead. See README.md.
//
// Reference: https://gobyexample.com/sorting-by-functions
package bylength

import "slices"

// ByLength sorts words in place, shortest first; words of equal length are
// in alphabetical order.
func ByLength(words []string) {
	slices.Sort(words)
}
