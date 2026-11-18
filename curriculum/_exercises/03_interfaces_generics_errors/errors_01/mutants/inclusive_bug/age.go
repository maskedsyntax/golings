// Mutant: rejects 150.
package age

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseAge(s string) (int, error) {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("age %q is not a whole number", s)
	}
	if n < 0 || n >= 150 {
		return 0, fmt.Errorf("age %d is out of range 0-150", n)
	}
	return n, nil
}
