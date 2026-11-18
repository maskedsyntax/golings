// Mutant: zero is reported as positive.
package classify

import (
	"strconv"
	"strings"
)

func Classify(s string) string {
	if n, err := strconv.Atoi(strings.TrimSpace(s)); err != nil {
		return "not a number"
	} else if n < 0 {
		return "negative"
	}
	return "positive"
}
