# panic_01: Must or panic

A **panic** means something went wrong that the program doesn't expect to handle. Most Go code
returns errors instead, but the standard library has a convention: a function named `Must…`
panics on bad input. Use it for values the programmer controls, such as a constant regular
expression.

```go
re := regexp.MustCompile(`^[a-z]+$`) // panics if the pattern is broken
```

## Your task

Implement `MustPort(s string) int`. It returns the port number in `s`, and it **panics** if
`s` is not a number between 1 and 65535. (`strconv.Atoi` converts a string to an int.)

Reference: https://gobyexample.com/panic
