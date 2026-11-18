# recover_02: Recover must be called directly

`Handle` should turn a panicking handler into an error, using a helper `capture` so the
recovery logic can be reused. Yet the panic still escapes.

The rule: `recover` stops a panic only when it is called **directly by the deferred
function**, not by a function that the deferred function calls.

## Your task

Fix `Handle` so the panic is recovered and reported as the error.

Reference: https://gobyexample.com/recover
