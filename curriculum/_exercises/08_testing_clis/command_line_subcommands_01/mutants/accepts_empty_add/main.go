// Mutant: add without text succeeds.
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
		fmt.Fprintln(stderr, "usage: todo <add|list> [flags]")
		return 2
	}
	switch args[0] {
	case "add":
		fs := flag.NewFlagSet("add", flag.ContinueOnError)
		fs.SetOutput(stderr)
		prio := fs.Int("p", 1, "priority")
		if err := fs.Parse(args[1:]); err != nil {
			return 2
		}
		fmt.Fprintf(stdout, "added: %s (priority %d)\n", strings.Join(fs.Args(), " "), *prio)
	case "list":
		fs := flag.NewFlagSet("list", flag.ContinueOnError)
		fs.SetOutput(stderr)
		all := fs.Bool("all", false, "")
		if err := fs.Parse(args[1:]); err != nil {
			return 2
		}
		if *all {
			fmt.Fprintln(stdout, "listing all tasks")
		} else {
			fmt.Fprintln(stdout, "listing open tasks")
		}
	default:
		fmt.Fprintf(stderr, "unknown subcommand %q\n", args[0])
		return 2
	}
	return 0
}
