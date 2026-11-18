// regular_expressions_03: Replacement templates (reference solution)
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
	// Braces end the group name; "$2_" would mean a group called "2_".
	return strings.ToLower(nameRE.ReplaceAllString(full, "${2}_${1}"))
}
