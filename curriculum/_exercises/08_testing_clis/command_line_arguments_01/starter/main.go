// command_line_arguments_01: os.Args
//
// The program greets itself. Fix it. See README.md.
//
// Reference: https://gobyexample.com/command-line-arguments
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

// run greets every name in args (the arguments, without the program name).
func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: greet NAME...")
		return 2
	}
	for _, name := range args {
		fmt.Fprintf(stdout, "Hello, %s!\n", name)
	}
	return 0
}
