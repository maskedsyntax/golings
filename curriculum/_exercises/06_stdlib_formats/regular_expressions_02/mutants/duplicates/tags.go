// Mutant: keeps repeated tags.
package tags

import (
	"regexp"
	"strings"
)

var tagRE = regexp.MustCompile(`#(\w+)`)

func HashTags(text string) []string {
	var tags []string
	for _, m := range tagRE.FindAllStringSubmatch(text, -1) {
		tags = append(tags, strings.ToLower(m[1]))
	}
	return tags
}
