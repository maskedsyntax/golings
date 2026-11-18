// string_functions_01: Slugs
//
// Slugify doesn't compile. Fix it. See README.md.
//
// Reference: https://gobyexample.com/string-functions
package slug

import "strings"

// Slugify lowercases title and joins its words with hyphens.
func Slugify(title string) string {
	words := strings.Fields(strings.Lower(title))
	return strings.Join("-", words)
}
