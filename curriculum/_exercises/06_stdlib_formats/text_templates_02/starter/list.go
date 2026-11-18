// text_templates_02: Lists with range
//
// Write the template. See README.md.
//
// Reference: https://gobyexample.com/text-templates
package list

import (
	"strings"
	"text/template"
)

// TODO: write the template.
const listTmpl = `TODO
`

var tmpl = template.Must(template.New("list").Parse(listTmpl))

// Render formats items as a shopping list.
func Render(items []string) (string, error) {
	var b strings.Builder
	if err := tmpl.Execute(&b, items); err != nil {
		return "", err
	}
	return b.String(), nil
}
