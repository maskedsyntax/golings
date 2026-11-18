// strings_and_runes_04: A Caesar cipher (reference solution)
package caesar

import "strings"

// Caesar shifts each ASCII letter of s by shift places, wrapping around the alphabet.
func Caesar(s string, shift int) string {
	k := rune(((shift % 26) + 26) % 26) // 0..25, even for negative shifts
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			r = 'a' + (r-'a'+k)%26
		case 'A' <= r && r <= 'Z':
			r = 'A' + (r-'A'+k)%26
		}
		b.WriteRune(r)
	}
	return b.String()
}
