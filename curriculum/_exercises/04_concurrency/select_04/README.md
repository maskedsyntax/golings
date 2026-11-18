# select_04: Sending to whoever is free

`select` cases can **send** as well as receive. A send case is ready when a receiver is
waiting (or the buffer has room):

```go
select {
case a <- v:
case b <- v:
}
```

## Your task

Implement `Balance(in <-chan int, a, b chan<- int, n int)`. It takes `n` values from `in`
and sends each one to **whichever** of `a` or `b` is ready to receive it. If one consumer
is stuck, every value goes to the other.

Reference: https://gobyexample.com/select
