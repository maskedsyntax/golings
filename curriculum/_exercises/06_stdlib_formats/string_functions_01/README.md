# string_functions_01: Slugs

The `strings` package has the everyday text helpers:

```go
strings.ToLower("Go")                    // "go"
strings.Fields("  a  b ")                // ["a" "b"]  (splits on any run of spaces)
strings.Join([]string{"a", "b"}, "-")    // "a-b"
```

A *slug* is a title turned into a URL-friendly name: `"  Hello  Go World "` becomes
`"hello-go-world"`.

## Your task

`Slugify` doesn't compile. Fix it.

Reference: https://gobyexample.com/string-functions
