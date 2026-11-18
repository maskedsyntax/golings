// regular_expressions_01: Anchors (reference solution)
package code

import "regexp"

// ^ and $ make the pattern cover the whole string.
var codeRE = regexp.MustCompile(`^[A-Z]{2}-\d{4}$`)

// IsValidCode reports whether s is exactly a product code like "AB-1234".
func IsValidCode(s string) bool {
	return codeRE.MatchString(s)
}
