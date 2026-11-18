# sorting_by_functions_03: Subtraction is not comparison

`Richest` sorts accounts by balance, largest first. It uses a well-known shortcut:
returning `b - a` from the comparison function. That works for small numbers and fails for
large ones.

## Your task

Find out why (the hints help) and fix `Richest`.

Reference: https://gobyexample.com/sorting-by-functions
