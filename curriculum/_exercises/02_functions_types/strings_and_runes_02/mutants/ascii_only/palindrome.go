// Mutant: only keeps ASCII letters and digits.
package palindrome

func IsPalindrome(s string) bool {
	var rs []rune
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			rs = append(rs, r)
		case r >= 'A' && r <= 'Z':
			rs = append(rs, r+'a'-'A')
		}
	}
	for i := range len(rs) / 2 {
		if rs[i] != rs[len(rs)-1-i] {
			return false
		}
	}
	return true
}
