# channels_02: Producer and consumer

Channels let one goroutine **produce** values while another **consumes** them.

## Your task

Implement two functions:

- `Generate(nums ...int) <-chan int` returns a channel and, from a new goroutine, sends each
  number on it in order.
- `Take(ch <-chan int, n int) []int` receives exactly `n` values from `ch` and returns them.

```go
Take(Generate(4, 5, 6), 2) // [4 5]
```

Don't close the channel yet. Closing channels has its own topic later in this module.

Reference: https://gobyexample.com/channels
