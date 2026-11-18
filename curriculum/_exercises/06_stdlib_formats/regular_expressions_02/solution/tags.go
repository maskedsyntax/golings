// regular_expressions_02: Hashtags (reference solution)
package tags

import (
	"regexp"
	"strings"
)

var tagRE = regexp.MustCompile(`#(\w+)`)

// HashTags returns the distinct hashtags in text, lowercased, in order of first appearance.
func HashTags(text string) []string {
	var tags []string
	seen := make(map[string]bool)
	for _, m := range tagRE.FindAllStringSubmatch(text, -1) {
		tag := strings.ToLower(m[1]) // m[0] is "#Go", m[1] is "Go"
		if !seen[tag] {
			seen[tag] = true
			tags = append(tags, tag)
		}
	}
	return tags
}
