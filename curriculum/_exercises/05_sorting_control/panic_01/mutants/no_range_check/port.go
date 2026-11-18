// Mutant: accepts any integer.
package port

import (
	"fmt"
	"strconv"
)

func MustPort(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("invalid port %q: %w", s, err))
	}
	return n
}
