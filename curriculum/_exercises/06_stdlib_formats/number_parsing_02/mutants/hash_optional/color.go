// Mutant: accepts colors without '#'.
package color

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseColor(s string) (r, g, b uint8, err error) {
	digits := strings.TrimPrefix(s, "#")
	if len(digits) == 3 {
		d := digits
		digits = string([]byte{d[0], d[0], d[1], d[1], d[2], d[2]})
	}
	if len(digits) != 6 {
		return 0, 0, 0, fmt.Errorf("bad color %q", s)
	}
	var parts [3]uint8
	for i := range parts {
		v, err := strconv.ParseUint(digits[2*i:2*i+2], 16, 8)
		if err != nil {
			return 0, 0, 0, err
		}
		parts[i] = uint8(v)
	}
	return parts[0], parts[1], parts[2], nil
}
