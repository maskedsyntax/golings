# worker_pools_03: Stopping on the first error

`ProcessAll` runs `f` on every item with a pool of workers. If any call fails, there is no
point processing the rest. `ProcessAll` must:

- return the **first** error (or `nil` if every call succeeded),
- stop handing out new items soon after that error (items already running may finish), and
- leave no goroutines behind when it returns.

## Your task

Implement `ProcessAll`. You can use a WaitGroup and channels; you don't need any other
`sync` types.

Reference: https://gobyexample.com/worker-pools
