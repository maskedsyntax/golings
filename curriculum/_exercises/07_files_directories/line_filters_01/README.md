# line_filters_01: An uppercase filter

A *line filter* reads text line by line, changes it, and writes it out, like `grep`, `sed`
or `tr`. Writing it against `io.Reader` and `io.Writer` (not `os.Stdin` and `os.Stdout`)
makes it easy to test:

```go
func main() {
	if err := Upper(os.Stdin, os.Stdout); err != nil { ... }
}
```

## Your task

Implement `Upper(r io.Reader, w io.Writer) error`. It writes every line of `r` in uppercase,
each ending with `\n` (even if the input's last line didn't). Return any read error.

Reference: https://gobyexample.com/line-filters
