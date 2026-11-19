// Mutant: cuts at the last dot anywhere in the path.
package paths

import (
	"path/filepath"
	"strings"
)

func ChangeExt(p, ext string) string {
	if i := strings.LastIndex(p, "."); i >= 0 {
		return p[:i] + ext
	}
	return p + ext
}

func RelTo(base, target string) (string, error) {
	return filepath.Rel(base, target)
}
