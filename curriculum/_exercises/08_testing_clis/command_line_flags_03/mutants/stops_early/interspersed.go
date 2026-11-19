// Mutant: plain Parse, stops at the first operand.
package interspersed

import "flag"

func ParseInterspersed(fs *flag.FlagSet, args []string) ([]string, error) {
	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	return fs.Args(), nil
}
