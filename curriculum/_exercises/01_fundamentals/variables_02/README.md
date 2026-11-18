# variables_02: The shadowed balance

This code **compiles** but always returns `0, false`. The culprit is *shadowing*: `:=` inside
an inner block declares a new variable with the same name as an outer one, so the outer one
never changes.

## Your task

Fix `Balance(transactions []int) (balance int, overdrawn bool)` so that:

- `balance` is the sum of all transactions,
- `overdrawn` is true if the running balance **ever** dropped below zero, even if it later recovered.

Reference: https://gobyexample.com/variables
