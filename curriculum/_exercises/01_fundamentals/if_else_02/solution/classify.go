// if_else_02: if with a short statement (reference solution)
package classify

import (
	"strconv"
	"strings"
)

// Classify describes the integer written in s.
func Classify(s string) string {
	if n, err := strconv.Atoi(strings.TrimSpace(s)); err != nil {
		return "not a number"
	} else if n < 0 {
		return "negative"
	} else if n == 0 {
		return "zero"
	}
	return "positive"
}
