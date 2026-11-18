// Mutant: Length is fixed, Reverse still reverses bytes.
package text

import "unicode/utf8"

func Length(s string) int { return utf8.RuneCountInString(s) }

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
