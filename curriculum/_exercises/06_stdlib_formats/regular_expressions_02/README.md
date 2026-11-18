# regular_expressions_02: Hashtags

Besides yes/no matching, a regexp can find every match and pull out *capture groups*:

```go
re := regexp.MustCompile(`(\w+)@(\w+)`)
re.FindAllStringSubmatch("a@b c@d", -1) // [["a@b" "a" "b"] ["c@d" "c" "d"]]
```

## Your task

Implement `HashTags(text string) []string`. It returns the tags in `text` (a `#` followed by
letters, digits or underscores):

- in lowercase, without the `#`,
- in the order they first appear,
- each tag only once.

```go
HashTags("Loving #Go and #rust_lang, #go again; # alone") // ["go" "rust_lang"]
```

Reference: https://gobyexample.com/regular-expressions
