// text_templates_01: Template fields (reference solution)
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

// Fields are looked up by their exported Go name.
var summary = template.Must(template.New("summary").Parse(
	"Hello, {{.Name}}! You ordered {{len .Items}} items."))

// Summary describes the order in one sentence.
func Summary(o Order) (string, error) {
	var b strings.Builder
	if err := summary.Execute(&b, o); err != nil {
		return "", err
	}
	return b.String(), nil
}
