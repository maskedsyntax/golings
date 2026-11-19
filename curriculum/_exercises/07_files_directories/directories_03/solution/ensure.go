// directories_03: Creating directories safely (reference solution)
package ensure

import (
	"fmt"
	"os"
	"path/filepath"
)

// EnsureDirs creates each relative directory path in rels under root.
func EnsureDirs(root string, rels []string) error {
	for _, rel := range rels {
		if !filepath.IsLocal(rel) { // rejects "..", absolute paths and ""
			return fmt.Errorf("%q is not a path inside %s", rel, root)
		}
		if err := os.MkdirAll(filepath.Join(root, rel), 0o755); err != nil {
			return err
		}
	}
	return nil
}
