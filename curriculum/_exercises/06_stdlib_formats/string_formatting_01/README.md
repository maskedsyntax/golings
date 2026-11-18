# string_formatting_01: Lining up columns

`fmt.Printf` takes a format string full of *verbs*: `%s` for strings, `%d` for integers,
`%f` for floats, `%v` for anything. Between the `%` and the verb you can add a width, a
precision and flags:

```go
fmt.Printf("|%6d|%-6d|\n", 42, 42)   // |    42|42    |
fmt.Printf("%.2f\n", 3.14159)         // 3.14
```

## Your task

Change only the format strings in `main.go` so the program prints exactly:

```text
item      |  qty|   price
apple     |    3|    0.50
banana    |   12|    0.25
"go" ff 00000101 +7 float64
```

The item column is 10 wide and left-aligned, `qty` is 5 wide, and `price` is 8 wide with two
decimal places.

Reference: https://gobyexample.com/string-formatting
