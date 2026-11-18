// string_functions_03: Whole words (reference solution)
package words

import (
	"strings"
	"unicode"
)

// CountWord reports how often word occurs in text as a whole word, ignoring case.
func CountWord(text, word string) int {
	notWordChar := func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }
	n := 0
	for _, w := range strings.FieldsFunc(text, notWordChar) {
		if strings.EqualFold(w, word) {
			n++
		}
	}
	return n
}
