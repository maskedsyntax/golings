# custom_errors_02: The typed nil error

This is the most famous gotcha with custom errors. A function whose result type is `error`
must return a literal `nil` for success. Returning a nil `*MyError` through the `error`
interface produces a **non-nil** error, because the interface holds a type.

## Your task

`Check(words)` should return `nil` when every word is lowercase, and a `*CheckError` listing
the offending words otherwise. Callers report a failure even for perfectly fine input. Fix
`Check`.

Reference: https://gobyexample.com/custom-errors
