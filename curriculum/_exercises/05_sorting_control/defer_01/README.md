# defer_01: Always clean up

`defer` schedules a function call to run when the surrounding function **returns**, on every
path, including early returns:

```go
f := createFile("/tmp/defer.txt")
defer closeFile(f)
writeFile(f)
```

`Copy` opens a source and a destination and closes them at the end. When a write fails it
returns early, and both stay open: a resource leak.

## Your task

Use `defer` (the `ast` check requires it) so `Copy` closes both resources on every path.

Reference: https://gobyexample.com/defer
