// sorting_01: A sorted copy
//
// Alphabetize changes the caller's slice. Fix it. See README.md.
//
// Reference: https://gobyexample.com/sorting
package alpha

import "slices"

// Alphabetize returns the words sorted alphabetically, leaving words unchanged.
func Alphabetize(words []string) []string {
	slices.Sort(words)
	return words
}
