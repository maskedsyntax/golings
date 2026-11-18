// Mutant: loads, but keeps the stray newlines.
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
{{range .}}
{{upper .Name}}: {{.Score}}
{{end}}
`

var tmpl = template.Must(template.New("scores").
	Funcs(template.FuncMap{"upper": strings.ToUpper}).
	Parse(scoresTmpl))

func Render(scores []Score) (string, error) {
	var b strings.Builder
	if err := tmpl.Execute(&b, scores); err != nil {
		return "", err
	}
	return b.String(), nil
}
