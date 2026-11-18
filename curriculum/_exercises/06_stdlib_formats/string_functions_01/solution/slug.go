// string_functions_01: Slugs (reference solution)
package slug

import "strings"

// Slugify lowercases title and joins its words with hyphens.
func Slugify(title string) string {
	words := strings.Fields(strings.ToLower(title))
	return strings.Join(words, "-")
}
