package pages

import (
	"slices"
	"testing"
)

func TestNames(t *testing.T) {
	if got, want := Names(), []string{"bye.tmpl", "hello.tmpl"}; !slices.Equal(got, want) {
		t.Fatalf("Names() = %q, want %q", got, want)
	}
}

func TestRender(t *testing.T) {
	for _, tt := range []struct{ name, data, want string }{
		{"hello.tmpl", "Ada", "Hello, Ada!\n"},
		{"bye.tmpl", "Bob", "Bye, Bob.\n"},
	} {
		got, err := Render(tt.name, tt.data)
		if err != nil || got != tt.want {
			t.Errorf("Render(%q, %q) = %q, %v; want %q, nil", tt.name, tt.data, got, err, tt.want)
		}
	}
	for _, name := range []string{"_draft.tmpl", "missing.tmpl"} {
		if _, err := Render(name, "x"); err == nil {
			t.Errorf("Render(%q) returned no error", name)
		}
	}
}
