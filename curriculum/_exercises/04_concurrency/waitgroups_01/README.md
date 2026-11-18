# waitgroups_01: Add before you go

A `sync.WaitGroup` waits for a collection of goroutines:

```go
var wg sync.WaitGroup
wg.Add(1)          // one more to wait for
go func() {
	defer wg.Done() // this one is finished
	work()
}()
wg.Wait()          // blocks until the counter is zero
```

`ProcessAll` sometimes returns before the work is done, and the race detector complains.

## Your task

Fix where `wg.Add` is called.

Reference: https://gobyexample.com/waitgroups
