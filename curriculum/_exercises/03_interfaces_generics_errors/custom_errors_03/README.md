# custom_errors_03: Unwrap

`fmt.Errorf("…%w…")` isn't the only way to wrap. A custom error type joins the chain by
providing an `Unwrap() error` method that returns its cause. Then `errors.Is` and
`errors.As` can see through it.

## Your task

1. Give `*QueryError` two methods:
   - `Error()` returns `query "<Query>": <cause's message>`.
   - `Unwrap()` returns the cause.
2. Implement `Run(q string) error`. It calls `exec(q)`, and if that fails, returns a
   `*QueryError` wrapping the failure. On success it returns `nil`.

Afterwards both of these hold for a failing query:

```go
errors.Is(err, ErrTimeout)    // the cause is still visible
errors.As(err, &queryErr)     // and so is the QueryError
```

Reference: https://gobyexample.com/custom-errors
