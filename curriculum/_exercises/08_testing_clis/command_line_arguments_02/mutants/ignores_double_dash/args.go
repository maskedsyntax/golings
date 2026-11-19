// Mutant: treats "--" as an ordinary flag.
package args

import "strings"

func Split(args []string) (flags, operands []string) {
	for _, a := range args {
		if len(a) > 1 && strings.HasPrefix(a, "-") {
			flags = append(flags, a)
		} else {
			operands = append(operands, a)
		}
	}
	return flags, operands
}
