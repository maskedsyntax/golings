// text_templates_01: Template fields
//
// The template fails to execute. Fix it. See README.md.
//
// Reference: https://gobyexample.com/text-templates
package order

import (
	"strings"
	"text/template"
)

// Order is what a customer bought.
type Order struct {
	Name  string
	Items []string
}

var summary = template.Must(template.New("summary").Parse(
	"Hello, {{.name}}! You ordered {{len .Items}} items."))

// Summary describes the order in one sentence.
func Summary(o Order) (string, error) {
	var b strings.Builder
	if err := summary.Execute(&b, o); err != nil {
		return "", err
	}
	return b.String(), nil
}
