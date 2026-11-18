# mutexes_02: The copied lock

`Counter` locks its mutex around every update, yet the count stays at zero, and `go vet`
complains.

## Your task

Fix `Counter` so `Inc` changes the counter it is called on and the mutex is never copied.
This exercise also runs `go vet`.

Reference: https://gobyexample.com/mutexes
