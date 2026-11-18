// regular_expressions_01: Anchors
//
// IsValidCode accepts strings that merely contain a code. Fix it. See README.md.
//
// Reference: https://gobyexample.com/regular-expressions
package code

import "regexp"

var codeRE = regexp.MustCompile(`[A-Z]{2}-\d{4}`)

// IsValidCode reports whether s is exactly a product code like "AB-1234".
func IsValidCode(s string) bool {
	return codeRE.MatchString(s)
}
