# waitgroups_02: Don't copy the WaitGroup

`RunWorkers` hangs. Each worker calls `wg.Done()`, so where does the count go?

Types from `sync` such as `WaitGroup` and `Mutex` must **never be copied** once in use. Passing
one by value makes a copy. `go vet` catches this mistake, and this exercise also runs it.

## Your task

Fix `RunWorkers` and `worker` so the workers count down the WaitGroup that `RunWorkers` is
waiting on.

Reference: https://gobyexample.com/waitgroups
