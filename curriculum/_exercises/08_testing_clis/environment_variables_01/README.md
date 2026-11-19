# environment_variables_01: Unset versus empty

Environment variables are a common way to configure programs. There's a subtle difference
between a variable that is **not set** and one that is set to the empty string:

```go
os.Getenv("PORT")            // "" in both cases
v, ok := os.LookupEnv("PORT") // ok tells them apart
```

## Your task

Fix `Port() (int, error)`:

| `PORT` | result |
|---|---|
| unset | `8080` |
| `"9000"` | `9000` |
| `""` (set but empty) | error: probably a mistake in the deployment |
| `"abc"`, `"0"`, `"70000"` | error |

Reference: https://gobyexample.com/environment-variables
