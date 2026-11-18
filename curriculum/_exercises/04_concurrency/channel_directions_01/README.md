# channel_directions_01: Send-only and receive-only

A function parameter can restrict a channel to one direction:

```go
func producer(out chan<- int) { out <- 1 }   // may only send
func consumer(in <-chan int)  { fmt.Println(<-in) } // may only receive
```

A bidirectional `chan int` converts automatically to either direction when you pass it.
The compiler then rejects code that uses the channel the wrong way.

The directions in this file are backwards, so it doesn't compile.

## Your task

Fix the parameter types of `Ping` and `Pong`. Every channel parameter must keep a direction
(no plain `chan string`); the `ast` check enforces this.

Reference: https://gobyexample.com/channel-directions
