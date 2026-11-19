// embed_directive_02: Embedding a directory (reference solution)
package pages

import (
	"embed"
	"io/fs"
	"strings"
	"text/template"
)

// Files starting with "_" or "." are excluded; "all:templates" would include them.
//
//go:embed templates
var files embed.FS

var tmpl = template.Must(template.ParseFS(files, "templates/*.tmpl"))

// Names returns the embedded template file names, sorted.
func Names() []string {
	entries, err := fs.ReadDir(files, "templates")
	if err != nil {
		panic(err) // the directory is compiled in; this cannot fail
	}
	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return names
}

// Render executes the named template with data.
func Render(name string, data any) (string, error) {
	var b strings.Builder
	if err := tmpl.ExecuteTemplate(&b, name, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
