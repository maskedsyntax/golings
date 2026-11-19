// Mutant: keeps the trailing newline.
package version

import _ "embed"

//go:embed version.txt
var raw string

func Version() string {
	return raw
}
