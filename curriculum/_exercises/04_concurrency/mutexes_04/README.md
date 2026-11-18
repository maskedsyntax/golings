# mutexes_04: Compute once per key

`Cache.Get` returns the cached value for a key, computing it on a miss. It uses a
`sync.RWMutex`, which allows many readers **or** one writer.

Computing a value is expensive, so each key must be computed **only once**, even when many
goroutines ask for the same missing key at the same moment. Right now they all compute it.

## Your task

Fix `Get` so `compute` runs at most once per key. Cache hits should still take only the read
lock.

Reference: https://gobyexample.com/mutexes
