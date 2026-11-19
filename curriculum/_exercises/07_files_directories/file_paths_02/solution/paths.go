// file_paths_02: Extensions and relative paths (reference solution)
package paths

import (
	"path/filepath"
	"strings"
)

// ChangeExt replaces the extension of the last element of p with ext.
func ChangeExt(p, ext string) string {
	return strings.TrimSuffix(p, filepath.Ext(p)) + ext
}

// RelTo returns target relative to base.
func RelTo(base, target string) (string, error) {
	return filepath.Rel(base, target)
}
