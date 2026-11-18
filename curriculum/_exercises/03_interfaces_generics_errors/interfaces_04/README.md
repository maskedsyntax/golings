# interfaces_04: Implementing io.Writer

Small interfaces are everywhere in Go's standard library. The most famous is
[`io.Writer`](https://pkg.go.dev/io#Writer):

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

Implement it once, and your type works with `fmt.Fprintf`, `io.Copy`, `bufio`, and much more.

## Your task

Give `*CountingWriter` a `Write` method so it satisfies `io.Writer`. It wraps another writer `W`:

- It forwards every write to `W`, and returns exactly what `W` returned.
- It adds the number of bytes `W` **actually** wrote to `N`. That can be fewer than `len(p)`
  if `W` fails partway.

Reference: https://gobyexample.com/interfaces
