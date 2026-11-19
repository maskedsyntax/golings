// embed_directive_01: Embedding a file
//
// Version returns "". Fix the embedding. See README.md.
//
// Reference: https://gobyexample.com/embed-directive
package version

import "strings"

// go:embed version.txt
var raw string

// Version returns the embedded version number.
func Version() string {
	return strings.TrimSpace(raw)
}
