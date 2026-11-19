package port

import (
	"os"
	"testing"
)

func TestPortUnset(t *testing.T) {
	t.Setenv("PORT", "") // restores the original value after the test
	os.Unsetenv("PORT")
	if p, err := Port(); err != nil || p != 8080 {
		t.Fatalf("PORT unset: Port() = %d, %v; want 8080, nil", p, err)
	}
}

func TestPortSet(t *testing.T) {
	t.Setenv("PORT", "9000")
	if p, err := Port(); err != nil || p != 9000 {
		t.Fatalf("PORT=9000: Port() = %d, %v; want 9000, nil", p, err)
	}
}

func TestPortInvalid(t *testing.T) {
	for _, v := range []string{"", "abc", "0", "70000", "-1"} {
		t.Setenv("PORT", v)
		if p, err := Port(); err == nil {
			t.Errorf("PORT=%q: Port() = %d, nil; want an error", v, p)
		}
	}
}
