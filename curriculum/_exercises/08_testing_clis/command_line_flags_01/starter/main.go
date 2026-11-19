// command_line_flags_01: Parse before you read
//
// The flags never take effect, and errors are ignored. Fix run. See README.md.
//
// Reference: https://gobyexample.com/command-line-flags
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

	greeting := fmt.Sprintf("Hello, %s!", *name)
	fs.Parse(args)

	for range *times {
		fmt.Fprintln(stdout, greeting)
	}
	return 0
}
