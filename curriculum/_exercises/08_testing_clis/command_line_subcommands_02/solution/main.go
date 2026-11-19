// command_line_subcommands_02: Flag sets that remember (reference solution)
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: tool <greet|count> [flags]")
		return 2
	}
	switch args[0] {
	case "greet":
		// A fresh FlagSet per call starts from the defaults every time.
		fs := flag.NewFlagSet("greet", flag.ContinueOnError)
		fs.SetOutput(stderr)
		loud := fs.Bool("loud", false, "shout")
		if err := fs.Parse(args[1:]); err != nil {
			return 2
		}
		for _, name := range fs.Args() {
			msg := "hello, " + name
			if *loud {
				msg = strings.ToUpper(msg)
			}
			fmt.Fprintln(stdout, msg)
		}
	case "count":
		fs := flag.NewFlagSet("count", flag.ContinueOnError)
		fs.SetOutput(stderr)
		step := fs.Int("by", 1, "step")
		if err := fs.Parse(args[1:]); err != nil {
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
