# command_line_flags_03: Flags after operands

Go's `flag` package stops at the first non-flag argument:

```text
tool a.txt -v      # -v is NOT parsed; fs.Args() is ["a.txt", "-v"]
```

Many users expect flags to work anywhere, as they do in most GNU tools.

## Your task

Implement `ParseInterspersed(fs *flag.FlagSet, args []string) ([]string, error)`. It parses
flags wherever they appear and returns the operands in order:

```go
ParseInterspersed(fs, []string{"a.txt", "-v", "b.txt", "-o", "out"})
// operands ["a.txt", "b.txt"]; -v and -o are set
```

- `--` still ends flag parsing: `["--", "a", "-v"]` gives operands `["a", "-v"]` and leaves
  `-v` unset.
- A lone `-` is an operand.
- Return any parse error.

Reference: https://gobyexample.com/command-line-flags
