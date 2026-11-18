package caesar

import "testing"

func TestCaesar(t *testing.T) {
	for _, tt := range []struct {
		s     string
		shift int
		want  string
	}{
		{"Hello, World!", 3, "Khoor, Zruog!"},
		{"xyz", 3, "abc"},
		{"XYZ", 3, "ABC"},
		{"abc", -1, "zab"},
		{"a", 27, "b"},
		{"a", -27, "z"},
		{"a", 26, "a"},
		{"née!", 1, "oéf!"},
		{"Go 1.24", 13, "Tb 1.24"},
		{"", 5, ""},
	} {
		if got := Caesar(tt.s, tt.shift); got != tt.want {
			t.Errorf("Caesar(%q, %d) = %q, want %q", tt.s, tt.shift, got, tt.want)
		}
	}
}

func TestCaesarRoundTrip(t *testing.T) {
	const s = "The quick brown fox jumps over the lazy dog."
	for _, k := range []int{1, 7, 25, 100, -3} {
		if got := Caesar(Caesar(s, k), -k); got != s {
			t.Errorf("Caesar(Caesar(s, %d), %d) = %q, want the original", k, -k, got)
		}
	}
}
