# constants_02: Human-readable sizes with iota

`iota` numbers the lines of a `const` block, and a line without an expression repeats the one
above it. That makes it easy to define a series such as KB, MB, GB.

## Your task

1. Declare package-level constants `KB`, `MB` and `GB` (1024, 1024², 1024³) using `iota`.
   *This exercise checks that the constants exist.*
2. Implement `HumanSize(n int64) string`:

| n | result |
|---|---|
| `0` | `"0 B"` |
| `1023` | `"1023 B"` |
| `1536` | `"1.5 KB"` |
| `1048576` | `"1.0 MB"` |
| `5 * GB` | `"5.0 GB"` |

Sizes of a terabyte or more are still shown in GB (`"2048.0 GB"`).

Reference: https://gobyexample.com/constants
