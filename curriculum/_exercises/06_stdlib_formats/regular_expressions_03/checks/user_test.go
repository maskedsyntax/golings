package user

import "testing"

func TestUsername(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{"Lovelace, Ada", "ada_lovelace"},
		{"  Hopper,Grace ", "grace_hopper"},
		{"Turing", ""},
		{"", ""},
	} {
		if got := Username(tt.in); got != tt.want {
			t.Errorf("Username(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
