# enums_01: iota and String

Go has no `enum` keyword. The idiom is a named integer type plus constants numbered by
`iota`, which counts 0, 1, 2… within a `const` block:

```go
type ServerState int
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
)
```

A `String()` method makes such values print nicely.

## Your task

The `Weekday` constants are defined in `weekday.go`. Implement:

- `String()` returns `"Sunday"` … `"Saturday"`. Any other value prints as `"Weekday(9)"`.
- `IsWeekend()` reports whether the day is Saturday or Sunday.
- `Next()` returns the following day, wrapping from Saturday to Sunday.

Reference: https://gobyexample.com/enums
