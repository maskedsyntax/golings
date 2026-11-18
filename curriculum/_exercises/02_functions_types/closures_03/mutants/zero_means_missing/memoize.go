// Mutant: treats a cached 0 as "not cached".
package memoize

func Memoize(f func(int) int) func(int) int {
	cache := map[int]int{}
	return func(n int) int {
		if v := cache[n]; v != 0 {
			return v
		}
		v := f(n)
		cache[n] = v
		return v
	}
}
