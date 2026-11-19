# context_03: Request-scoped values

Contexts also carry **request-scoped values**, such as a request ID for logging:

```go
ctx = context.WithValue(ctx, key, value)
v := ctx.Value(key)
```

This package has two bugs:

1. It stores the request ID under the string key `"requestID"`, so any other code using that
   string can overwrite or fake it.
2. `ForAudit` is for work that must keep going after the request ends (writing an audit log),
   so it must not be cancelled with the request. It returns `context.Background()`, which also
   loses the request ID.

## Your task

Fix both. `RequestID` must see only IDs set with `WithRequestID`. `ForAudit` must keep the
values and ignore the parent's cancellation.

Reference: https://gobyexample.com/context
