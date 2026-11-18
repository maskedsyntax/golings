// strings_and_runes_02: Unicode-aware palindromes (reference solution)
package palindrome

import "unicode"

// IsPalindrome reports whether the letters and digits of s, ignoring case,
// read the same in both directions.
func IsPalindrome(s string) bool {
	var rs []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			rs = append(rs, unicode.ToLower(r))
		}
	}
	for i := range len(rs) / 2 {
		if rs[i] != rs[len(rs)-1-i] {
			return false
		}
	}
	return true
}
