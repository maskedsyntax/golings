// closures_01: A counter that remembers (reference solution)
package counter

// Counter returns a function that returns 1, 2, 3, ... on successive calls.
func Counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
