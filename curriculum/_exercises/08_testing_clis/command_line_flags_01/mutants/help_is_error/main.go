// Mutant: treats -h as a usage error.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(stderr)
	name := fs.String("name", "World", "who to greet")
	times := fs.Int("n", 1, "how many times to greet")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	for range *times {
		fmt.Fprintf(stdout, "Hello, %s!\n", *name)
	}
	return 0
}
