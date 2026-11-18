// Mutant: compares words case-sensitively.
package words

import (
	"strings"
	"unicode"
)

func CountWord(text, word string) int {
	n := 0
	for _, w := range strings.FieldsFunc(text, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }) {
		if w == word {
			n++
		}
	}
	return n
}
