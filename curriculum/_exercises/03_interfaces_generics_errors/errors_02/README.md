# errors_02: Sentinel errors and wrapping

A **sentinel error** is a predeclared error value that callers can check for:

```go
var ErrOutOfTea = errors.New("no more tea available")
```

To add context without losing the sentinel, **wrap** it with `%w`. Then `errors.Is`
recognizes it anywhere in the chain:

```go
err := fmt.Errorf("making tea: %w", ErrOutOfTea)
errors.Is(err, ErrOutOfTea) // true
```

## Your task

Implement `Price(prices map[string]int, key string) (int, error)`:

- It returns the price and `nil` when `key` exists with a non-negative price.
- A missing key is an error that **wraps `ErrNotFound`**, and whose message includes the key.
- A negative price is an error that **wraps `ErrInvalid`**, and whose message includes the key.

Reference: https://gobyexample.com/errors
