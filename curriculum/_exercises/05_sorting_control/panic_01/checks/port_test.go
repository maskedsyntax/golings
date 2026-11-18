package port

import "testing"

func panics(f func()) (didPanic bool) {
	defer func() {
		if recover() != nil {
			didPanic = true
		}
	}()
	f()
	return false
}

func TestMustPortValid(t *testing.T) {
	for _, tt := range []struct {
		in   string
		want int
	}{{"80", 80}, {"8080", 8080}, {"1", 1}, {"65535", 65535}} {
		var got int
		if panics(func() { got = MustPort(tt.in) }) {
			t.Errorf("MustPort(%q) panicked on a valid port", tt.in)
		} else if got != tt.want {
			t.Errorf("MustPort(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestMustPortPanics(t *testing.T) {
	for _, in := range []string{"", "http", "0", "65536", "-1", "80a"} {
		if !panics(func() { MustPort(in) }) {
			t.Errorf("MustPort(%q) did not panic", in)
		}
	}
}
