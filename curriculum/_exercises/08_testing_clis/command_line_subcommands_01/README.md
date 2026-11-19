# command_line_subcommands_01: A tool with subcommands

Tools like `go` and `git` have **subcommands** (`go build`, `go test`), each with its own
flags. In Go, each subcommand gets its own `flag.FlagSet`.

## Your task

Implement `run(args, stdout, stderr) int` for a small `todo` tool:

| command | output (stdout) |
|---|---|
| `todo add [-p N] TEXT...` | `added: TEXT (priority N)`, where `-p` defaults to 1 and the words of TEXT are joined with spaces |
| `todo list [-all]` | `listing all tasks` with `-all`, otherwise `listing open tasks` |

These all print a message to stderr and return 2: no subcommand, an unknown subcommand (name it
in the message), invalid flags, and `add` without text. Success returns 0.

Reference: https://gobyexample.com/command-line-subcommands
