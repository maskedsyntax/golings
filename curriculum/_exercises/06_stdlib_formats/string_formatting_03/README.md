# string_formatting_03: A String method that calls itself

A type with a `String() string` method controls how `%v` and `%s` print it:

```go
type Celsius float64
func (c Celsius) String() string { ... }

fmt.Println(Celsius(21.5)) // 21.5°C
```

This package's `String` method formats `c` with `%v`... which calls `String` again, forever.
`go test` refuses to run it, because `go vet` spots the problem first.

## Your task

Fix `String` so that a temperature prints with one decimal place, like `21.5°C` or `-3.0°C`.
`Report` already uses it to print aligned columns.

Reference: https://gobyexample.com/string-formatting
