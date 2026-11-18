// Mutant: groups in the wrong order.
package user

import (
	"regexp"
	"strings"
)

var nameRE = regexp.MustCompile(`^\s*(\w+),\s*(\w+)\s*$`)

func Username(full string) string {
	if !nameRE.MatchString(full) {
		return ""
	}
	return strings.ToLower(nameRE.ReplaceAllString(full, "${1}_${2}"))
}
