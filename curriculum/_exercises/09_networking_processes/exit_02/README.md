# exit_02: Choosing exit codes

The exit status is how a program tells scripts what happened. A common convention:

| status | meaning |
|---|---|
| 0 | success |
| 1 | general failure |
| 2 | usage error (bad flags or arguments) |
| 3 | something wasn't found (this program's own code) |
| 130 | interrupted (128 + SIGINT) |

## Your task

Implement `ExitCode(err error) int`, which `main` would use as `os.Exit(ExitCode(run()))`:

- `nil` → 0
- a `*UsageError` anywhere in the chain → 2
- `ErrNotFound` anywhere in the chain → 3
- `context.Canceled` anywhere in the chain → 130
- anything else → 1

Errors are usually wrapped by the time they reach `main`, as in
`fmt.Errorf("loading %s: %w", name, ErrNotFound)`, and must still be recognized.

Reference: https://gobyexample.com/exit
