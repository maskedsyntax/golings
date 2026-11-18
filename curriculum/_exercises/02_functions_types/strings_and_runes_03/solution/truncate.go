// strings_and_runes_03: Truncate without breaking characters (reference solution)
package truncate

// Truncate shortens s to at most max characters, ending with "…" if it cut anything.
func Truncate(s string, max int) string {
	if max <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max-1]) + "…"
}
