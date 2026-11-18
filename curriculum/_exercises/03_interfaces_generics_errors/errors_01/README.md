# errors_01: Returning errors

Go reports failure with an explicit `error` result, by convention the last one:

```go
func f(arg int) (int, error) {
	if arg == 42 {
		return 0, errors.New("can't work with 42")
	}
	return arg + 3, nil
}
```

Callers check `if err != nil` immediately. On failure, the other results should be zero
values.

## Your task

Implement `ParseAge(s string) (int, error)`:

- Surrounding whitespace is ignored: `ParseAge(" 42 ")` → `42, nil`.
- Text that isn't a whole number is an error.
- Numbers outside 0–150 are an error.
- Whenever the error is non-nil, the returned age is `0`. Error messages should mention the
  bad input.

Reference: https://gobyexample.com/errors
