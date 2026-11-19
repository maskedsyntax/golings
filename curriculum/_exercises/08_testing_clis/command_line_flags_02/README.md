# command_line_flags_02: Custom flag types

Besides `String`, `Int` and `Bool`, the `flag` package has `Duration`, and it accepts **any type**
that implements `flag.Value`:

```go
type Value interface {
	String() string
	Set(string) error
}
```

## Your task

Make `TagList` a `flag.Value` and implement `Parse(args []string) (Config, error)`:

| flag | meaning |
|---|---|
| `-tag x` | may be repeated: `-tag a -tag b` gives `["a", "b"]`; an empty tag is an error |
| `-timeout 2m30s` | a `time.Duration`, default `5s` |
| `-v` | verbose, default `false` |

The remaining arguments are `Files`. Invalid input is an error. Send the flag package's own
messages to `io.Discard`.

Reference: https://gobyexample.com/command-line-flags
