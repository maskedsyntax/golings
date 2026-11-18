# constants_01: Typed and untyped constants

Go constants can be **untyped**: `const Factor = 2` is just "the number 2" and works with
an `int` or a `float64`. Give a constant an explicit type and it follows that type's rules,
including its range.

## Your task

Fix `constants.go` so it compiles:

- `Scale(x float64) float64` returns `x * Factor`.
- `TB` holds the number of bytes in a terabyte (2⁴⁰), usable as an `int64`.

Reference: https://gobyexample.com/constants
