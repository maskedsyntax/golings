# errors_03: Checking wrapped errors

Once an error is wrapped, `err == ErrNotFound` is false: `err` is a new value that
*contains* `ErrNotFound`. Use `errors.Is`, which walks the chain of wrapped errors.

## Your task

`GetOrDefault(key, def)` calls `Fetch(key)`:

- If the key isn't found, it returns `def` and a nil error.
- Any **other** error is returned to the caller unchanged.
- Otherwise it returns the fetched value.

Currently a missing key comes back as an error, and other errors are lost. Fix it.

Reference: https://gobyexample.com/errors
