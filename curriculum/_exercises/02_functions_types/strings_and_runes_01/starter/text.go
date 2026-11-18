// strings_and_runes_01: Bytes are not characters
//
// These work for ASCII but not for "héllo" or "日本". Fix them.
//
//	Length("héllo") == 5
//	Reverse("añb") == "bña"
//
// Reference: https://gobyexample.com/strings-and-runes
package text

// Length returns the number of characters (runes) in s.
func Length(s string) int {
	return len(s)
}

// Reverse returns s with its characters in reverse order.
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
