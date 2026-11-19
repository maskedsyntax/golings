# context_02: First answer wins

A common pattern is to ask several replicas the same question, take the first answer, and
**cancel the rest** so they stop wasting work.

## Your task

Implement `FirstResult(ctx, fns...) (string, error)`:

- Run every function concurrently, giving each a context.
- Return the first **successful** result, and cancel the context the others received.
- If every function fails, return all their errors, joined so `errors.Is` matches each one.
- If `ctx` itself is cancelled first, return `ctx.Err()`.
- With no functions at all, return an error.

Reference: https://gobyexample.com/context
