// command_line_subcommands_02: Flag sets that remember
//
// Flags leak between calls, and no arguments panics. Fix it. See README.md.
//
// Reference: https://gobyexample.com/command-line-subcommands
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	greetCmd = flag.NewFlagSet("greet", flag.ContinueOnError)
	loud     = greetCmd.Bool("loud", false, "shout")

	countCmd = flag.NewFlagSet("count", flag.ContinueOnError)
	step     = countCmd.Int("by", 1, "step")
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	switch args[0] {
	case "greet":
		greetCmd.SetOutput(stderr)
		if err := greetCmd.Parse(args[1:]); err != nil {
			return 2
		}
		for _, name := range greetCmd.Args() {
			msg := "hello, " + name
			if *loud {
				msg = strings.ToUpper(msg)
			}
			fmt.Fprintln(stdout, msg)
		}
	case "count":
		countCmd.SetOutput(stderr)
		if err := countCmd.Parse(args[1:]); err != nil {
			return 2
		}
		for i := 0; i < 3; i++ {
			fmt.Fprintln(stdout, i**step)
		}
	default:
		fmt.Fprintf(stderr, "unknown subcommand %q\n", args[0])
		return 2
	}
	return 0
}
