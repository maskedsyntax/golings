// Mutant: a package-level variable is shared by every counter.
package counter

var n int

func Counter() func() int {
	return func() int {
		n++
		return n
	}
}
