# json_02: Strict config with defaults

`json.Unmarshal` decodes JSON into a Go value. Fields whose keys are absent keep whatever value
they already had, and unknown keys are silently ignored. That second part is a problem for
config files: a typo like `"prot": 9000` would quietly do nothing.

## Your task

Implement `ParseConfig(data []byte) (Config, error)`:

- Missing `host` means `"localhost"` and missing `port` means `8080`.
- An unknown key is an error.
- Invalid JSON, or a value of the wrong type, is an error.
- `port` must be between 1 and 65535.

Reference: https://gobyexample.com/json
