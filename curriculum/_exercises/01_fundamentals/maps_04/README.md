# maps_04: Writing to a nil map

`var m map[string]int` gives you a **nil map**. Reading from it is fine (you get zero values),
but writing to it panics.

## Your task

`Merge(a, b map[string]int) map[string]int` should return a **new** map with every key from
`a` and `b`, summing the values of keys that appear in both. It currently panics. Fix it,
and make sure `a` and `b` are left unchanged.

Reference: https://gobyexample.com/maps
