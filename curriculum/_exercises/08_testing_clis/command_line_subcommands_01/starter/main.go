// command_line_subcommands_01: A tool with subcommands
//
// Implement run. See README.md.
//
// Reference: https://gobyexample.com/command-line-subcommands
package main

import (
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	// TODO
	return 0
}
