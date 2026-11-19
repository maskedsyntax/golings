# context_01: Cancellable waits

A `context.Context` carries a cancellation signal across API boundaries. `ctx.Done()` is a
channel that is closed when the context is cancelled or its deadline passes, and `ctx.Err()`
says why (`context.Canceled` or `context.DeadlineExceeded`).

`Wait(ctx, d)` should pause for `d`, but give up as soon as `ctx` is done. It ignores `ctx`
completely.

## Your task

Fix `Wait`. Return `nil` after the full wait, or `ctx.Err()` if the context ends first.

Reference: https://gobyexample.com/context
