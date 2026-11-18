# methods_03: Methods on named types

You can define methods on any named type declared in your package, not only on structs:

```go
type Celsius float64
func (c Celsius) Fahrenheit() float64 { … }
```

A type with a `String() string` method satisfies `fmt.Stringer`, and the `fmt` package uses
that method when printing the value.

## Your task

In `temperature.go`, implement:

- `(c Celsius) Fahrenheit() float64`, which converts with `°F = °C × 9/5 + 32`.
- `(c Celsius) String() string`, which formats to one decimal place: `"21.5°C"`, `"-3.0°C"`.
  `fmt.Sprint(Celsius(21.5))` must use it.
- `(r Readings) Warmest() (Celsius, bool)`, which returns the highest reading. It returns
  `0, false` when there are none.

Reference: https://gobyexample.com/methods
