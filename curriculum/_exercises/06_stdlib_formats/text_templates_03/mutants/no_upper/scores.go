// Mutant: right whitespace, but names aren't capitalized.
package scores

import (
	"strings"
	"text/template"
)

type Score struct {
	Name  string
	Score int
}

const scoresTmpl = `
{{- range .}}
{{- .Name}}: {{.Score}}
{{end -}}
`

var tmpl = template.Must(template.New("scores").Parse(scoresTmpl))

func Render(scores []Score) (string, error) {
	var b strings.Builder
	if err := tmpl.Execute(&b, scores); err != nil {
		return "", err
	}
	return b.String(), nil
}
