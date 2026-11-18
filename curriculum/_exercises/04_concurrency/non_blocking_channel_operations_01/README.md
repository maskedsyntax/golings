# non_blocking_channel_operations_01: Try, don't wait

Plain sends and receives block. A `select` with a `default` case doesn't:

```go
select {
case msg := <-messages:
	fmt.Println("received", msg)
default:
	fmt.Println("no message waiting")
}
```

## Your task

Implement two helpers that never block:

- `TrySend(ch chan<- string, msg string) bool` sends if possible and reports whether it did.
- `TryReceive(ch <-chan string) (string, bool)` receives if a value is ready right now.

Reference: https://gobyexample.com/non-blocking-channel-operations
