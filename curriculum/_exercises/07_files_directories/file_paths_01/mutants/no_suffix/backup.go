// Mutant: forgets the .bak suffix.
package backup

import "path/filepath"

func BackupPath(p string) string {
	return filepath.Join(filepath.Dir(p), "backup", filepath.Base(p))
}
