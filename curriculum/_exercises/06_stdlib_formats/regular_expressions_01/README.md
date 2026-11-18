# regular_expressions_01: Anchors

`regexp` compiles a pattern once and then matches it against strings:

```go
var re = regexp.MustCompile(`p([a-z]+)ch`)
re.MatchString("peach") // true
```

A product code is exactly two capital letters, a hyphen and four digits, like `AB-1234`.
`IsValidCode` accepts `AB-1234`, but it also accepts `xAB-1234` and `AB-12345`.

## Your task

Fix the pattern so the **whole** string has to be a code.

Reference: https://gobyexample.com/regular-expressions
