# select_02: Until told to stop

A common pattern is a loop that keeps receiving work until a **quit** channel says stop:

```go
for {
	select {
	case v := <-data:
		// handle v
	case <-quit:
		return
	}
}
```

## Your task

Implement `Collect(data <-chan int, quit <-chan struct{}) []int`. It receives values from
`data` until `quit` is closed, then returns everything it received, in order.

Reference: https://gobyexample.com/select
