// Mutant: avoids the recursion but prints as many decimals as %v would.
package temp

import (
	"fmt"
	"strings"
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%v°C", float64(c))
}

type Reading struct {
	City string
	Temp Celsius
}

func Report(rs []Reading) string {
	var b strings.Builder
	for _, r := range rs {
		fmt.Fprintf(&b, "%-6s%7s\n", r.City, r.Temp)
	}
	return b.String()
}
