# strings_and_runes_04: A Caesar cipher

A rune is an integer (an alias for `int32`), so `'c' - 'a' == 2` and `'a' + 2 == 'c'`.
`strings.Builder` builds a string piece by piece without repeated copying.

## Your task

Implement `Caesar(s string, shift int) string`. It shifts each ASCII letter `shift` places
along the alphabet, wrapping around from `z` to `a`.

- It keeps the case: `Caesar("Hello", 3) == "Khoor"`.
- `shift` may be negative or larger than 26: `Caesar("a", -1) == "z"`, `Caesar("a", 27) == "b"`.
- Everything else, including digits, punctuation and non-ASCII letters, stays unchanged:
  `Caesar("née!", 1) == "oéf!"`.

Reference: https://gobyexample.com/strings-and-runes
