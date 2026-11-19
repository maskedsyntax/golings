// environment_variables_01: Unset versus empty (reference solution)
package port

import (
	"fmt"
	"os"
	"strconv"
)

// Port returns the TCP port from $PORT, or 8080 if it is unset.
func Port() (int, error) {
	v, ok := os.LookupEnv("PORT")
	if !ok {
		return 8080, nil
	}
	p, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("PORT=%q: %w", v, err)
	}
	if p < 1 || p > 65535 {
		return 0, fmt.Errorf("PORT=%d is out of range", p)
	}
	return p, nil
}
