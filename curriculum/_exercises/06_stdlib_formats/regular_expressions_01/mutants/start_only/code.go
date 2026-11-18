// Mutant: anchors only the start.
package code

import "regexp"

var codeRE = regexp.MustCompile(`^[A-Z]{2}-\d{4}`)

func IsValidCode(s string) bool {
	return codeRE.MatchString(s)
}
