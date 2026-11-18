// Mutant: copies but forgets to sort.
package alpha

import "slices"

func Alphabetize(words []string) []string {
	return slices.Clone(words)
}
