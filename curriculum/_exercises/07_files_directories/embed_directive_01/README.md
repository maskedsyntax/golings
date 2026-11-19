# embed_directive_01: Embedding a file

The `//go:embed` directive copies files into your program at **compile time**, so the binary
doesn't need them at run time:

```go
import _ "embed"

//go:embed hello.txt
var greeting string
```

This exercise comes with `version.txt`. `Version` should return its contents without the
trailing newline, but it returns `""`.

## Your task

Fix the embedding. Look closely at the directive.

Reference: https://gobyexample.com/embed-directive
