// text_templates_02: Lists with range (reference solution)
package list

import (
	"strings"
	"text/template"
)

const listTmpl = `{{if .}}Shopping list:
{{range .}}- {{.}}
{{end}}{{else}}Nothing to buy.
{{end}}`

var tmpl = template.Must(template.New("list").Parse(listTmpl))

// Render formats items as a shopping list.
func Render(items []string) (string, error) {
	var b strings.Builder
	if err := tmpl.Execute(&b, items); err != nil {
		return "", err
	}
	return b.String(), nil
}
