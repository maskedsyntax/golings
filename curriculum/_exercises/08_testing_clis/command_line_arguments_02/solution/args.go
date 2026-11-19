// command_line_arguments_02: Flags, operands and -- (reference solution)
package args

import "strings"

// Split separates command-line arguments into flags and operands.
func Split(args []string) (flags, operands []string) {
	endOfFlags := false
	for _, a := range args {
		switch {
		case endOfFlags:
			operands = append(operands, a)
		case a == "--":
			endOfFlags = true
		case len(a) > 1 && strings.HasPrefix(a, "-"):
			flags = append(flags, a)
		default:
			operands = append(operands, a)
		}
	}
	return flags, operands
}
