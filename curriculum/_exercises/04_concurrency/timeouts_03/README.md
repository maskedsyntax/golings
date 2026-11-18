# timeouts_03: A timeout for each call

`CallAll` runs a batch of calls, giving **each** one at most `d`. It works, but it runs the
calls one after another, so four hung calls cost four timeouts.

## Your task

Make `CallAll` run every call concurrently, so the whole batch takes about `d` no matter how
many calls hang. `results[i]` must still describe `calls[i]`.

Reference: https://gobyexample.com/timeouts
