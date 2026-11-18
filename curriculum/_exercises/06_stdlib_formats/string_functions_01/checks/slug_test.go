package slug

import "testing"

func TestSlugify(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{"  Hello  Go World ", "hello-go-world"},
		{"Go", "go"},
		{"already-a-slug", "already-a-slug"},
		{"", ""},
	} {
		if got := Slugify(tt.in); got != tt.want {
			t.Errorf("Slugify(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
