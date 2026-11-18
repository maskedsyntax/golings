// Mutant: uses unicode.IsLetter, so non-ASCII letters get shifted too.
package caesar

import (
	"strings"
	"unicode"
)

func Caesar(s string, shift int) string {
	k := rune(((shift % 26) + 26) % 26)
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				r = 'A' + (r-'A'+k)%26
			} else {
				r = 'a' + (r-'a'+k)%26
			}
		}
		b.WriteRune(r)
	}
	return b.String()
}
