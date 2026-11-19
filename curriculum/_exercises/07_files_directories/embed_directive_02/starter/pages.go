// embed_directive_02: Embedding a directory
//
// Embed the templates directory and implement Names and Render. See README.md.
//
// Reference: https://gobyexample.com/embed-directive
package pages

import "errors"

// Names returns the embedded template file names, sorted.
func Names() []string {
	// TODO
	return nil
}

// Render executes the named template with data.
func Render(name string, data any) (string, error) {
	// TODO
	return "", errors.New("not implemented")
}
