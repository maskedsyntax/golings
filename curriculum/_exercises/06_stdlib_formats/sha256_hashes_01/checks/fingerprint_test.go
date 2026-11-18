package fingerprint

import "testing"

func TestFingerprint(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"abc", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
	} {
		if got := Fingerprint([]byte(tt.in)); got != tt.want {
			t.Errorf("Fingerprint(%q) =\n  %s\nwant\n  %s", tt.in, got, tt.want)
		}
	}
	if Fingerprint([]byte("abd")) == Fingerprint([]byte("abc")) {
		t.Error("different inputs gave the same fingerprint")
	}
}
