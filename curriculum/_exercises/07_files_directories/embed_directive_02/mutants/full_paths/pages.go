// Mutant: returns paths instead of names.
package pages

import (
	"embed"
	"io/fs"
	"strings"
	"text/template"
)

//go:embed templates
var files embed.FS

var tmpl = template.Must(template.ParseFS(files, "templates/*.tmpl"))

func Names() []string {
	names, _ := fs.Glob(files, "templates/*")
	return names
}

func Render(name string, data any) (string, error) {
	var b strings.Builder
	if err := tmpl.ExecuteTemplate(&b, name, data); err != nil {
		return "", err
	}
	return b.String(), nil
}
