# text_templates_03: Functions and whitespace

Templates can call your own functions once you register them with `Funcs`:

```go
t := template.New("x").Funcs(template.FuncMap{"upper": strings.ToUpper})
```

Text outside `{{ }}` is copied to the output exactly, newlines included. A minus sign next to
the braces trims the whitespace beside an action: `{{- .X}}` trims before it, `{{.X -}}` trims
after it.

`Render` should print one score per line, with the name in capitals:

```text
ADA: 9
BOB: 7
```

## Your task

Fix the package so it loads, then fix the template's whitespace so the output matches exactly.
Keep the template on several lines; that's why the trim markers exist.

Reference: https://gobyexample.com/text-templates
