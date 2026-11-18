# pointers_01: Changing the caller's variable

Go passes every argument **by value**: a function receives a copy. To let a function change
the caller's variable, pass a pointer to it:

```go
func zero(p *int) { *p = 0 }
x := 5
zero(&x) // x is now 0
```

`&x` takes the address of `x`; `*p` reads or writes the value `p` points to.

## Your task

The tests call `Double(&x)` and `Swap(&a, &b)`. Fix `pointers.go` so that the code compiles and:

- `Double` doubles the caller's integer,
- `Swap` exchanges the values of the two integers.

Reference: https://gobyexample.com/pointers
