// Mutant: forgets to lowercase.
package slug

import "strings"

func Slugify(title string) string {
	return strings.Join(strings.Fields(title), "-")
}
