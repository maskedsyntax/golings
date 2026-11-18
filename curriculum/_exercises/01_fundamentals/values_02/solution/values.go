// values_02: Booleans and operators (reference solution)
package values

// CanVote reports whether someone may vote: aged 18 or over AND registered.
func CanVote(age int, registered bool) bool {
	return age >= 18 && registered
}

// Differ reports whether exactly one of a and b is true.
func Differ(a, b bool) bool {
	return a != b
}

// IsEven reports whether n is even (including negative n).
func IsEven(n int) bool {
	return n%2 == 0
}
