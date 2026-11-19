// Mutant: LookupEnv, but any integer is accepted.
package port

import (
	"os"
	"strconv"
)

func Port() (int, error) {
	v, ok := os.LookupEnv("PORT")
	if !ok {
		return 8080, nil
	}
	return strconv.Atoi(v)
}
