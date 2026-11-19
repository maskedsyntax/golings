package paths

import (
	"path/filepath"
	"testing"
)

var fs = filepath.FromSlash

func TestChangeExt(t *testing.T) {
	for _, tt := range []struct{ p, ext, want string }{
		{fs("a/b.txt"), ".md", fs("a/b.md")},
		{"archive.tar.gz", ".zip", "archive.tar.zip"},
		{fs("a.b/c"), ".txt", fs("a.b/c.txt")},
		{"README", ".md", "README.md"},
	} {
		if got := ChangeExt(tt.p, tt.ext); got != tt.want {
			t.Errorf("ChangeExt(%q, %q) = %q, want %q", tt.p, tt.ext, got, tt.want)
		}
	}
}

func TestRelTo(t *testing.T) {
	for _, tt := range []struct{ base, target, want string }{
		{fs("/srv/app"), fs("/srv/app/static/css"), fs("static/css")},
		{fs("/srv/app"), fs("/srv/data/x"), fs("../data/x")},
		{fs("/srv/app"), fs("/srv/app"), "."},
		{fs("a/b"), fs("a/c"), fs("../c")},
	} {
		got, err := RelTo(tt.base, tt.target)
		if err != nil || got != tt.want {
			t.Errorf("RelTo(%q, %q) = %q, %v; want %q, nil", tt.base, tt.target, got, err, tt.want)
		}
	}
	if got, err := RelTo("app", fs("/etc")); err == nil {
		t.Errorf("RelTo(\"app\", %q) = %q, nil; want an error", fs("/etc"), got)
	}
}
