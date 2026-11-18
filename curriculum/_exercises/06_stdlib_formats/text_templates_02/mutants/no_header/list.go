// Mutant: range/else without the header line.
package list

import (
	"strings"
	"text/template"
)

const listTmpl = `{{range .}}- {{.}}
{{else}}Nothing to buy.
{{end}}`

var tmpl = template.Must(template.New("list").Parse(listTmpl))

func Render(items []string) (string, error) {
	var b strings.Builder
	if err := tmpl.Execute(&b, items); err != nil {
		return "", err
	}
	return b.String(), nil
}
