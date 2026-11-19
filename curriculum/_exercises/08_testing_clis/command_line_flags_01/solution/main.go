// command_line_flags_01: Parse before you read (reference solution)
package main

import (
	"errors"
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
		if errors.Is(err, flag.ErrHelp) {
			return 0 // usage was already printed
		}
		return 2
	}

	greeting := fmt.Sprintf("Hello, %s!", *name) // read after Parse
	for range *times {
		fmt.Fprintln(stdout, greeting)
	}
	return 0
}
