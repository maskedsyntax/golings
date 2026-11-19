// Mutant: "all:" also embeds the draft.
package pages

import (
	"embed"
	"io/fs"
	"strings"
	"text/template"
)

//go:embed all:templates
var files embed.FS

var tmpl = template.Must(template.ParseFS(files, "templates/*.tmpl"))

func Names() []string {
	entries, _ := fs.ReadDir(files, "templates")
	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return names
}

func Render(name string, data any) (string, error) {
	var b strings.Builder
	if err := tmpl.ExecuteTemplate(&b, name, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
