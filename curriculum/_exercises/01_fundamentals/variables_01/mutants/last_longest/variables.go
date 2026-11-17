// Mutant: >= keeps the last of several equally long items.
package variables

func Describe(items []string) (count int, longest string) {
	for _, item := range items {
		count++
		if len(item) >= len(longest) {
			longest = item
		}
	}
	return count, longest
}
