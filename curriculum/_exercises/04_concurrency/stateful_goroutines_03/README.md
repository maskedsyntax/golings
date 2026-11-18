# stateful_goroutines_03: Shutting the owner down

A goroutine that owns state must eventually stop, or it leaks. `Counter` has a `Close` method,
but after `Close`, `Inc` hangs forever instead of reporting an error, and `Close` doesn't wait
for the owner to exit.

## Your task

Fix `Counter` so that:

- `Close` stops the owner goroutine and returns only once it has exited;
- `Inc` after `Close` returns `0, ErrClosed` immediately instead of blocking;
- `Close` is called at most once (you don't need to handle double closes).

`done` is already in the struct for you to use.

Reference: https://gobyexample.com/stateful-goroutines
