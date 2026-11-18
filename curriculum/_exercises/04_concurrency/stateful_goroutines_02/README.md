# stateful_goroutines_02: The missing reply

This store's owner goroutine handles reads, writes and deletes. Calling `Delete` hangs forever.

## Your task

Fix the store so `Delete` returns once the key is gone.

Reference: https://gobyexample.com/stateful-goroutines
