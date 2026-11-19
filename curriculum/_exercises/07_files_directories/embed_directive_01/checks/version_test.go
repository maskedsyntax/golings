package version

import "testing"

func TestVersion(t *testing.T) {
	if got := Version(); got != "1.4.2" {
		t.Fatalf("Version() = %q, want \"1.4.2\"", got)
	}
}
