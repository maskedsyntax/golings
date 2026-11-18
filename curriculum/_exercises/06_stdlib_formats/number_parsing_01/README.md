# number_parsing_01: Atoi returns two values

`strconv` converts strings to numbers. Every parse function also returns an error, because the
text might not be a number:

```go
n, err := strconv.Atoi("42")  // 42, nil
_, err = strconv.Atoi("x")    // err: strconv.Atoi: parsing "x": invalid syntax
```

## Your task

`Sum` adds up the space-separated integers in a string. It doesn't compile. Fix it, and make
it return the error when a field isn't a number.

Reference: https://gobyexample.com/number-parsing
