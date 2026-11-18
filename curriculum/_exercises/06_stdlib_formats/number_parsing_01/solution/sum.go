// number_parsing_01: Atoi returns two values (reference solution)
package sum

import (
	"strconv"
	"strings"
)

// Sum adds the space-separated integers in s.
func Sum(s string) (int, error) {
	total := 0
	for _, f := range strings.Fields(s) {
		n, err := strconv.Atoi(f)
		if err != nil {
			return 0, err
		}
		total += n
	}
	return total, nil
}
