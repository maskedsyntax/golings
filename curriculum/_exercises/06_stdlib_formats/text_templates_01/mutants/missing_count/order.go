// Mutant: drops the item count.
package order

import (
	"strings"
	"text/template"
)

type Order struct {
	Name  string
	Items []string
}

var summary = template.Must(template.New("summary").Parse("Hello, {{.Name}}! You ordered items."))

func Summary(o Order) (string, error) {
	var b strings.Builder
	if err := summary.Execute(&b, o); err != nil {
		return "", err
	}
	return b.String(), nil
}
