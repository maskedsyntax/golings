// sorting_01: A sorted copy (reference solution)
package alpha

import "slices"

// Alphabetize returns the words sorted alphabetically, leaving words unchanged.
func Alphabetize(words []string) []string {
	sorted := slices.Clone(words) // a new backing array
	slices.Sort(sorted)
	return sorted
}
