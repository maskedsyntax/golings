// text_templates_03: Functions and whitespace
//
// The package panics at load time, and the output has stray blank lines.
// Fix both. See README.md.
//
// Reference: https://gobyexample.com/text-templates
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

const scoresTmpl = `
{{range .}}
{{upper .Name}}: {{.Score}}
{{end}}
`

var tmpl = template.Must(template.New("scores").Parse(scoresTmpl)).
	Funcs(template.FuncMap{"upper": strings.ToUpper})

// Render prints one "NAME: score" line per score.
func Render(scores []Score) (string, error) {
	var b strings.Builder
	if err := tmpl.Execute(&b, scores); err != nil {
		return "", err
	}
	return b.String(), nil
}
