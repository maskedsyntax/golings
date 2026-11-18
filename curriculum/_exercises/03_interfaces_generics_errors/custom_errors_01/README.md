# custom_errors_01: An error type with fields

`error` is an interface with one method, `Error() string`, so any type can be an error.
A custom error type can carry structured details, which callers retrieve with `errors.As`:

```go
var ae *argError
if errors.As(err, &ae) {
	fmt.Println(ae.arg, ae.message)
}
```

## Your task

1. Give `*ValidationError` an `Error()` method returning `"<field>: <reason>"`, for example
   `"email: must contain @"`.
2. Implement `Validate(u User) error`. It checks the fields in this order and returns a
   `*ValidationError` for the **first** problem:
   - `Name` is empty → field `"name"`, reason `"must not be empty"`
   - `Email` has no `@` → field `"email"`, reason `"must contain @"`
   - `Age` is negative → field `"age"`, reason `"must not be negative"`

   It returns `nil` if the user is valid.

Reference: https://gobyexample.com/custom-errors
