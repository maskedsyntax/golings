// embed_directive_01: Embedding a file (reference solution)
package version

import (
	_ "embed" // required by //go:embed
	"strings"
)

//go:embed version.txt
var raw string

// Version returns the embedded version number.
func Version() string {
	return strings.TrimSpace(raw)
}
