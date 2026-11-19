// directories_03: Creating directories safely
//
// EnsureDirs fails on nested paths and trusts its input. Fix it. See README.md.
//
// Reference: https://gobyexample.com/directories
package ensure

import (
	"os"
	"path/filepath"
)

// EnsureDirs creates each relative directory path in rels under root.
func EnsureDirs(root string, rels []string) error {
	for _, rel := range rels {
		if err := os.Mkdir(filepath.Join(root, rel), 0o755); err != nil {
			return err
		}
	}
	return nil
}
