# text_templates_02: Lists with range

Templates can loop and branch:

```text
{{range .}}...{{end}}          repeat for each element; inside, {{.}} is the element
{{if .}}...{{else}}...{{end}}  an empty slice counts as false
```

## Your task

Write `listTmpl` so `Render` produces, for `[]string{"apple", "bread"}`:

```text
Shopping list:
- apple
- bread
```

and for an empty list:

```text
Nothing to buy.
```

Each line, including the last, ends with a newline.

Reference: https://gobyexample.com/text-templates
