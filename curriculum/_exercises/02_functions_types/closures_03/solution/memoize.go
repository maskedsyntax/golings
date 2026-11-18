// closures_03: Memoize a function (reference solution)
package memoize

// Memoize returns a function equivalent to f that calls f at most once per argument.
func Memoize(f func(int) int) func(int) int {
	cache := map[int]int{}
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		v := f(n)
		cache[n] = v
		return v
	}
}
