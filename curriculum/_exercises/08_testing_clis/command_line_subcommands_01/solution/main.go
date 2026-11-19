// command_line_subcommands_01: A tool with subcommands (reference solution)
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

const usage = "usage: todo <add|list> [flags]"

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, usage)
		return 2
	}
	switch cmd, rest := args[0], args[1:]; cmd {
	case "add":
		fs := flag.NewFlagSet("add", flag.ContinueOnError)
		fs.SetOutput(stderr)
		prio := fs.Int("p", 1, "priority")
		if err := fs.Parse(rest); err != nil {
			return 2
		}
		text := strings.Join(fs.Args(), " ")
		if text == "" {
			fmt.Fprintln(stderr, "add: missing task text")
			return 2
		}
		fmt.Fprintf(stdout, "added: %s (priority %d)\n", text, *prio)
	case "list":
		fs := flag.NewFlagSet("list", flag.ContinueOnError)
		fs.SetOutput(stderr)
		all := fs.Bool("all", false, "include finished tasks")
		if err := fs.Parse(rest); err != nil {
			return 2
		}
		if *all {
			fmt.Fprintln(stdout, "listing all tasks")
		} else {
			fmt.Fprintln(stdout, "listing open tasks")
		}
	default:
		fmt.Fprintf(stderr, "unknown subcommand %q\n%s\n", cmd, usage)
		return 2
	}
	return 0
}
