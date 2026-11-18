package decode

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestDecodeAny(t *testing.T) {
	data := []byte{0xfb, 0xff, 0xfe, 'g', 'o', '?', '>'}
	for _, enc := range []struct {
		name string
		e    *base64.Encoding
	}{
		{"StdEncoding", base64.StdEncoding},
		{"RawStdEncoding", base64.RawStdEncoding},
		{"URLEncoding", base64.URLEncoding},
		{"RawURLEncoding", base64.RawURLEncoding},
	} {
		s := enc.e.EncodeToString(data)
		got, err := DecodeAny(s)
		if err != nil || !bytes.Equal(got, data) {
			t.Errorf("DecodeAny(%q) (%s) = %x, %v; want %x", s, enc.name, got, err, data)
		}
	}
	for _, s := range []string{"aGk=", "aGk"} {
		if got, err := DecodeAny(s); err != nil || string(got) != "hi" {
			t.Errorf("DecodeAny(%q) = %q, %v; want \"hi\"", s, got, err)
		}
	}
}

func TestDecodeAnyErrors(t *testing.T) {
	for _, s := range []string{"!!!!", "a", "aGk*"} {
		if _, err := DecodeAny(s); err == nil {
			t.Errorf("DecodeAny(%q) returned no error", s)
		}
	}
}
