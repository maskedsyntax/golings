// Mutant: forgets to fold case.
package palindrome

import "unicode"

func IsPalindrome(s string) bool {
	var rs []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			rs = append(rs, r)
		}
	}
	for i := range len(rs) / 2 {
		if rs[i] != rs[len(rs)-1-i] {
			return false
		}
	}
	return true
}
