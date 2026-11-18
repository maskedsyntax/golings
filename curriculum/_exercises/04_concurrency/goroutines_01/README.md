# goroutines_01: Starting goroutines

A **goroutine** is a function running concurrently with the rest of the program. Start one by
putting `go` in front of a call:

```go
go doWork()        // returns immediately; doWork runs in the background
```

`RunAll` should run every task **at the same time** and return once all of them have
finished. The waiting part (a `sync.WaitGroup`) is already written: you'll study it properly in
the waitgroups topic. Right now the tasks run one after another.

## Your task

Make `RunAll` start each task in its own goroutine.

The tests use tasks that wait for each other, so they only finish if they all run at once.

Reference: https://gobyexample.com/goroutines
