// file_paths_01: Portable paths
//
// BackupPath splits paths by hand. Fix it. See README.md.
//
// Reference: https://gobyexample.com/file-paths
package backup

import "strings"

// BackupPath returns where a backup of the file at p goes: a "backup"
// directory next to it, with ".bak" appended to the name.
func BackupPath(p string) string {
	i := strings.LastIndex(p, "/")
	return p[:i] + "/backup/" + p[i+1:] + ".bak"
}
