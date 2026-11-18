// closures_01: A counter that remembers
//
// Implement Counter. See README.md.
//
//	next := Counter()
//	next() == 1; next() == 2; next() == 3
//
// Reference: https://gobyexample.com/closures
package counter

// Counter returns a function that returns 1, 2, 3, ... on successive calls.
func Counter() func() int {
	// TODO
	return func() int { return 0 }
}
