// closures_03: Memoize a function
//
// Implement Memoize. See README.md.
//
//	square := Memoize(func(n int) int { return n * n })
//	square(4) == 16 // calls f
//	square(4) == 16 // answered from the cache
//
// Reference: https://gobyexample.com/closures
package memoize

// Memoize returns a function equivalent to f that calls f at most once per argument.
func Memoize(f func(int) int) func(int) int {
	// TODO
	return f
}
