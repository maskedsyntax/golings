# variadic_functions_02: At least one argument

A variadic parameter must come last, but it can follow ordinary parameters. That's a neat
way to require at least one argument:

```go
func Max(first int, rest ...int) int
```

`Max()` is now a compile error instead of a runtime problem.

## Your task

Implement, in `variadic.go`:

- `Max(first int, rest ...int) int` returns the largest of its arguments.
- `JoinNonEmpty(sep string, parts ...string) string` joins the parts that aren't empty
  strings, with `sep` between them: `JoinNonEmpty(", ", "a", "", "b")` → `"a, b"`.

Reference: https://gobyexample.com/variadic-functions
