# variables_01: Declare, assign, don't shadow

`var x int` and `x := 0` both declare variables, and `x = 1` assigns to one that
already exists. Mixing them up is one of the most common early Go mistakes.

## Your task

Fix `variables.go` so it compiles and `Describe(items)` returns:

- `count`: the number of items.
- `longest`: the longest item. On a tie, keep the **first** one seen. With no items, return `""`.

Reference: https://gobyexample.com/variables
