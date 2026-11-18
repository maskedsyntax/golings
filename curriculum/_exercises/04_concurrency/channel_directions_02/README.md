# channel_directions_02: A pipeline stage

A **pipeline** chains stages together with channels. Each stage receives from an input
channel and sends results on an output channel:

```go
squares := Square(Square(numbers)) // fourth powers
```

## Your task

Implement `Square(in <-chan int) <-chan int`. It returns a channel immediately, and a goroutine
receives each value from `in` and sends its square on the returned channel.

The input never closes in this exercise (closing channels comes later), so the goroutine can
loop forever. The `ast` check requires the parameter to stay receive-only.

Reference: https://gobyexample.com/channel-directions
