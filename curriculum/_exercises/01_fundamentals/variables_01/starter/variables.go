// variables_01: Declare, assign, don't shadow
//
// This file does not compile. Fix it so that
//
//	Describe([]string{"go", "gopher", "fun"}) returns (3, "gopher")
//
// and ties keep the first longest item.
//
// Reference: https://gobyexample.com/variables
package variables

// Describe counts items and finds the longest one (first wins on ties).
func Describe(items []string) (count int, longest string) {
	count := 0
	var total int

	for _, item := range items {
		count++
		if len(item) > len(longest) {
			longest := item
		}
	}
	return count, longest
}
