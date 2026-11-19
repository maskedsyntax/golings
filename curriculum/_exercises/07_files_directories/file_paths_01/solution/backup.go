// file_paths_01: Portable paths (reference solution)
package backup

import "path/filepath"

// BackupPath returns where a backup of the file at p goes: a "backup"
// directory next to it, with ".bak" appended to the name.
func BackupPath(p string) string {
	return filepath.Join(filepath.Dir(p), "backup", filepath.Base(p)+".bak")
}
