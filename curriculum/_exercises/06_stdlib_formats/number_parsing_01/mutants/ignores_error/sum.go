// Mutant: discards the error.
package sum

import (
	"strconv"
	"strings"
)

func Sum(s string) (int, error) {
	total := 0
	for _, f := range strings.Fields(s) {
		n, _ := strconv.Atoi(f)
		total += n
	}
	return total, nil
}
