// regular_expressions_03: Replacement templates
//
// Username loses the first name. Fix it. See README.md.
//
// Reference: https://gobyexample.com/regular-expressions
package user

import (
	"regexp"
	"strings"
)

var nameRE = regexp.MustCompile(`^\s*(\w+),\s*(\w+)\s*$`)

// Username turns "Last, First" into "first_last". It returns "" for any other input.
func Username(full string) string {
	if !nameRE.MatchString(full) {
		return ""
	}
	return strings.ToLower(nameRE.ReplaceAllString(full, "$2_$1"))
}
