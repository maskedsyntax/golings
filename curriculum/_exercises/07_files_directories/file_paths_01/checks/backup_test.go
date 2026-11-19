package backup

import (
	"path/filepath"
	"testing"
)

func TestBackupPath(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{filepath.Join("docs", "notes.txt"), filepath.Join("docs", "backup", "notes.txt.bak")},
		{filepath.Join("a", "b", "c.tar.gz"), filepath.Join("a", "b", "backup", "c.tar.gz.bak")},
		{"notes.txt", filepath.Join("backup", "notes.txt.bak")},
	} {
		if got := BackupPath(tt.in); got != tt.want {
			t.Errorf("BackupPath(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
