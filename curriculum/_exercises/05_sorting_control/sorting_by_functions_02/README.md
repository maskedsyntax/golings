# sorting_by_functions_02: Sorting records

`slices.SortFunc` works on slices of structs too. The comparison function decides which
fields matter, and in what order.

## Your task

Implement `SortPeople(people []Person)`. It sorts in place by `Age` (youngest first), then by
`Name` (alphabetically) for people of the same age.

Reference: https://gobyexample.com/sorting-by-functions
