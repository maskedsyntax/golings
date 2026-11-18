// strings_and_runes_01: Bytes are not characters (reference solution)
package text

import "unicode/utf8"

// Length returns the number of characters (runes) in s.
func Length(s string) int {
	return utf8.RuneCountInString(s)
}

// Reverse returns s with its characters in reverse order.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
