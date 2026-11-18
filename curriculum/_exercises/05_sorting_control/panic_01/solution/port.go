// panic_01: Must or panic (reference solution)
package port

import (
	"fmt"
	"strconv"
)

// MustPort returns the port number in s. It panics if s is not a valid
// port (an integer from 1 to 65535).
func MustPort(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("invalid port %q: %w", s, err))
	}
	if n < 1 || n > 65535 {
		panic(fmt.Errorf("port %d out of range 1-65535", n))
	}
	return n
}
