// Mutant: MkdirAll, but no IsLocal check.
package ensure

import (
	"os"
	"path/filepath"
)

func EnsureDirs(root string, rels []string) error {
	for _, rel := range rels {
		if err := os.MkdirAll(filepath.Join(root, rel), 0o755); err != nil {
			return err
		}
	}
	return nil
}
