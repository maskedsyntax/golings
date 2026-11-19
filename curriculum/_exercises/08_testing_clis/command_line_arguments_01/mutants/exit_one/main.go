// Mutant: right arguments, wrong usage status.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: greet NAME...")
		return 1
	}
	for _, name := range args {
		fmt.Fprintf(stdout, "Hello, %s!\n", name)
	}
	return 0
}
