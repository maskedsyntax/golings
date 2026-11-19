# command_line_flags_01: Parse before you read

The `flag` package defines flags that return **pointers**. Their values are filled in by
`Parse`:

```go
fs := flag.NewFlagSet("greet", flag.ContinueOnError)
name := fs.String("name", "World", "who to greet")
err := fs.Parse(args)   // *name changes here
```

With `ContinueOnError`, `Parse` returns errors instead of exiting the program, which keeps
`run` testable.

## Your task

Fix `run`:

- `greet -name Ada -n 2` prints `Hello, Ada!` twice and returns 0.
- `greet` prints `Hello, World!` once.
- `greet -h` prints only the usage (to stderr) and returns 0.
- An unknown flag prints the error and usage to stderr, writes nothing to stdout, and returns 2.

Reference: https://gobyexample.com/command-line-flags
