// environment_variables_01: Unset versus empty
//
// Port can't tell unset from empty, and accepts any number. Fix it. See README.md.
//
// Reference: https://gobyexample.com/environment-variables
package port

import (
	"os"
	"strconv"
)

// Port returns the TCP port from $PORT, or 8080 if it is unset.
func Port() (int, error) {
	v := os.Getenv("PORT")
	if v == "" {
		return 8080, nil
	}
	return strconv.Atoi(v)
}
