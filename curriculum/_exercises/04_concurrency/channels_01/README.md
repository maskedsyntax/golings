# channels_01: Sending on a channel

**Channels** connect goroutines. One goroutine sends a value with `ch <- v` and another
receives it with `v := <-ch`. On an unbuffered channel (`make(chan T)`) the send **blocks
until a receiver is ready**, and the receive blocks until a sender is ready.

`Greet` should return a channel that will deliver one greeting:

```go
msg := <-Greet("Gopher") // "Hello, Gopher!"
```

Right now `Greet` never returns.

## Your task

Fix `Greet`.

Reference: https://gobyexample.com/channels
