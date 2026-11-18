# custom_errors_04: Reporting every problem

A form with three mistakes should report all three, not make the user fix them one at a
time. [`errors.Join`](https://pkg.go.dev/errors#Join) combines several errors into one:

- it returns `nil` if all its arguments are nil,
- its message puts each error on its own line,
- `errors.Is` and `errors.As` match any of the joined errors.

## Your task

Implement `ValidateAll(u User) error`. It checks the same three rules as `custom_errors_01`,
but reports **every** problem, each as a `*FieldError`, joined with `errors.Join`, in the
order name, email, age. It returns `nil` if there are no problems.

```go
err := ValidateAll(User{Email: "x", Age: -1})
fmt.Println(err)
// name: must not be empty
// email: must contain @
// age: must not be negative
```

Reference: https://gobyexample.com/custom-errors
