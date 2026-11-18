// string_formatting_03: A String method that calls itself (reference solution)
package temp

import (
	"fmt"
	"strings"
)

// Celsius is a temperature in degrees Celsius.
type Celsius float64

// String formats c with one decimal place, e.g. "21.5°C".
func (c Celsius) String() string {
	// float64(c) has no String method, so this doesn't recurse.
	return fmt.Sprintf("%.1f°C", float64(c))
}

// Reading is a temperature measured in a city.
type Reading struct {
	City string
	Temp Celsius
}

// Report prints one reading per line, the city left-aligned in 6 columns and the
// temperature right-aligned in 7.
func Report(rs []Reading) string {
	var b strings.Builder
	for _, r := range rs {
		fmt.Fprintf(&b, "%-6s%7s\n", r.City, r.Temp)
	}
	return b.String()
}
