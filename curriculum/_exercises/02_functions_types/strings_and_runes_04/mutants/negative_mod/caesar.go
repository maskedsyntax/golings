// Mutant: % on a negative shift gives a negative offset.
package caesar

import "strings"

func Caesar(s string, shift int) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			r = 'a' + (r-'a'+rune(shift))%26
		case 'A' <= r && r <= 'Z':
			r = 'A' + (r-'A'+rune(shift))%26
		}
		b.WriteRune(r)
	}
	return b.String()
}
