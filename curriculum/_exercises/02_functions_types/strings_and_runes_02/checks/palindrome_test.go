package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, tt := range []struct {
		s    string
		want bool
	}{
		{"", true},
		{"a", true},
		{"racecar", true},
		{"Racecar", true},
		{"A man, a plan, a canal: Panama", true},
		{"No 'x' in Nixon", true},
		{"12321", true},
		{"hello", false},
		{"ab", false},
		{"Été", true},
		{"Ésope reste ici et se repose", false}, // É and the final e differ
		{"étè", false},
		{"日本日", true},
		{"日本", false},
	} {
		if got := IsPalindrome(tt.s); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
