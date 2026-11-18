# number_parsing_02: Hex colors

`strconv.ParseInt` and `ParseUint` take a **base** (2 to 36) and a **bit size** that the result
has to fit in:

```go
strconv.ParseUint("ff", 16, 8)  // 255
strconv.ParseUint("1ff", 16, 8) // error: value out of range
```

## Your task

Implement `ParseColor(s string) (r, g, b uint8, err error)` for CSS-style colors:

- `"#ff8800"` gives 255, 136, 0.
- The short form `"#f80"` means `"#ff8800"`.
- Upper- and lowercase hex digits both work.
- A missing `#`, the wrong number of digits, or a character that isn't a hex digit is an error.

Reference: https://gobyexample.com/number-parsing
