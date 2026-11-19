# command_line_subcommands_02: Flag sets that remember

`run` is called once per command, and tests call it many times in the same process, as a
long-running shell or server would.

```go
run([]string{"greet", "-loud", "ada"}, ...)  // HELLO, ADA
run([]string{"greet", "bob"}, ...)           // should be: hello, bob   (it shouts)
run(nil, ...)                                // should be: usage + 2    (it panics)
```

## Your task

Fix `run` so each call is independent, and so a missing subcommand is a usage error (message on
stderr, return 2).

Reference: https://gobyexample.com/command-line-subcommands
