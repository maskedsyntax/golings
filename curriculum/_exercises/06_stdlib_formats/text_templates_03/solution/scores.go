// text_templates_03: Functions and whitespace (reference solution)
package scores

import (
	"strings"
	"text/template"
)

// Score is one player's result.
type Score struct {
	Name  string
	Score int
}

// {{- removes the newline before an action; -}} removes the one after it.
const scoresTmpl = `
{{- range .}}
{{- upper .Name}}: {{.Score}}
{{end -}}
`

// Funcs must come before Parse, which needs to know every function name.
var tmpl = template.Must(template.New("scores").
	Funcs(template.FuncMap{"upper": strings.ToUpper}).
	Parse(scoresTmpl))

// Render prints one "NAME: score" line per score.
func Render(scores []Score) (string, error) {
	var b strings.Builder
	if err := tmpl.Execute(&b, scores); err != nil {
		return "", err
	}
	return b.String(), nil
}
