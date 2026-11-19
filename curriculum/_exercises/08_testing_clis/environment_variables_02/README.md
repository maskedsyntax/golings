# environment_variables_02: Prefixed settings

Programs often read all their settings from variables that share a prefix, like `APP_NAME` and
`APP_DB_URL`. `os.Environ()` lists every variable as a `KEY=value` string.

## Your task

Implement `Load(environ []string, prefix string) map[string]string`. Real code would call it as
`Load(os.Environ(), "APP_")`:

- Keep only variables whose name starts with `prefix`, and remove the prefix.
- Lowercase the rest of the name: `APP_DB_URL` becomes `db_url`.
- Split at the **first** `=`, because values may contain `=`.
- Skip malformed entries (no `=`) and names that are empty after removing the prefix.
- Always return a non-nil map.

Reference: https://gobyexample.com/environment-variables
