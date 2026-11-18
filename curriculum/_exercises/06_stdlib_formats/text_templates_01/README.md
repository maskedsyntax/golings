# text_templates_01: Template fields

`text/template` fills placeholders in text with data. Inside `{{ }}`, a dot means "the data",
and `.Field` reads a field of it:

```go
t := template.Must(template.New("hi").Parse("Hi, {{.Name}}!\n"))
t.Execute(os.Stdout, struct{ Name string }{"Ada"}) // Hi, Ada!
```

`Summary` should produce:

```text
Hello, Ada! You ordered 2 items.
```

but `Execute` returns an error instead.

## Your task

Fix the template.

Reference: https://gobyexample.com/text-templates
