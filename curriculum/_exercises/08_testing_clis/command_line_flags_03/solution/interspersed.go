// command_line_flags_03: Flags after operands (reference solution)
package interspersed

import "flag"

// ParseInterspersed parses flags anywhere in args and returns the operands.
func ParseInterspersed(fs *flag.FlagSet, args []string) ([]string, error) {
	var operands []string
	for {
		if err := fs.Parse(args); err != nil {
			return nil, err
		}
		rest := fs.Args()
		if len(rest) == 0 {
			return operands, nil
		}
		// Parse stops either at an operand or right after "--".
		if used := len(args) - len(rest); used > 0 && args[used-1] == "--" {
			return append(operands, rest...), nil
		}
		operands = append(operands, rest[0])
		args = rest[1:]
	}
}
