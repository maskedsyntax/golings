// number_parsing_01: Atoi returns two values
//
// Sum doesn't compile. Fix it. See README.md.
//
// Reference: https://gobyexample.com/number-parsing
package sum

import (
	"strconv"
	"strings"
)

// Sum adds the space-separated integers in s.
func Sum(s string) (int, error) {
	total := 0
	for _, f := range strings.Fields(s) {
		n := strconv.Atoi(f)
		total += n
	}
	return total, nil
}
