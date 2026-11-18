// Mutant: one package-level cache is shared by every memoized function.
package memoize

var cache = map[int]int{}

func Memoize(f func(int) int) func(int) int {
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		v := f(n)
		cache[n] = v
		return v
	}
}
