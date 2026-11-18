// string_functions_03: Whole words
//
// CountWord counts substrings, not words. Fix it. See README.md.
//
// Reference: https://gobyexample.com/string-functions
package words

import "strings"

// CountWord reports how often word occurs in text as a whole word, ignoring case.
func CountWord(text, word string) int {
	return strings.Count(strings.ToLower(text), strings.ToLower(word))
}
