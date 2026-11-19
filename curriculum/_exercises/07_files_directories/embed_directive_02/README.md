# embed_directive_02: Embedding a directory

`//go:embed` can embed whole directories into an `embed.FS`, which is a read-only file system
compiled into your program. Many standard functions accept it, such as `fs.ReadDir`,
`fs.ReadFile` and `template.ParseFS`.

This exercise comes with a `templates/` directory. It holds `hello.tmpl`, `bye.tmpl` and a
work-in-progress `_draft.tmpl`.

## Your task

Embed `templates` and implement:

- `Names() []string`: the embedded template file names, sorted:
  `["bye.tmpl", "hello.tmpl"]`. When you embed a directory, files whose names start with `_`
  or `.` are left out, so the draft isn't there.
- `Render(name string, data any) (string, error)`: executes the named template.
  `Render("hello.tmpl", "Ada")` gives `"Hello, Ada!\n"`. An unknown name is an error.

Reference: https://gobyexample.com/embed-directive
