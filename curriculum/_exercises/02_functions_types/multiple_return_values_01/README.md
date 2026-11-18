# multiple_return_values_01: Receiving two results

Go functions can return more than one value:

```go
func vals() (int, int) { return 3, 7 }
a, b := vals()
_, c := vals() // the blank identifier discards a value you don't need
```

## Your task

`divmod.go` doesn't compile. Fix it so that:

- `DivMod(a, b)` returns both the quotient and the remainder of `a / b`
  (`DivMod(7, 2)` → `3, 1`),
- `Quotient(a, b)` returns just the quotient, using `DivMod`.

Reference: https://gobyexample.com/multiple-return-values
