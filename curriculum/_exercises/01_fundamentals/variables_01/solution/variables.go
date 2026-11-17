// variables_01: Declare, assign, don't shadow (reference solution)
package variables

// Describe counts items and finds the longest one (first wins on ties).
func Describe(items []string) (count int, longest string) {
	count = 0
	for _, item := range items {
		count++
		if len(item) > len(longest) {
			longest = item
		}
	}
	return count, longest
}
