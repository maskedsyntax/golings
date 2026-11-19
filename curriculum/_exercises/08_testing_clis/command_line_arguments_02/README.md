# command_line_arguments_02: Flags, operands and --

Command-line programs follow a few conventions:

- Arguments starting with `-` are flags: `-v`, `--name=x`.
- A lone `-` is an ordinary operand. It usually means "standard input".
- `--` ends the flags. Everything after it is an operand, even `-x`. That's how
  `rm -- -weird-file` works.

## Your task

Implement `Split(args []string) (flags, operands []string)`, keeping the original order in
each list:

```go
Split([]string{"-v", "a.txt", "--name=x", "-", "--", "-not-a-flag", "b.txt"})
// flags:    ["-v", "--name=x"]
// operands: ["a.txt", "-", "-not-a-flag", "b.txt"]
```

Reference: https://gobyexample.com/command-line-arguments
