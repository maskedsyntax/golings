package token

import (
	"bytes"
	"strings"
	"testing"
)

func TestEncodeToken(t *testing.T) {
	for _, tt := range []struct {
		in   []byte
		want string
	}{
		{[]byte{0xfb, 0xff, 0xfe, 0xff}, "-__-_w"},
		{[]byte("hi"), "aGk"},
		{nil, ""},
	} {
		got := EncodeToken(tt.in)
		if got != tt.want {
			t.Errorf("EncodeToken(%x) = %q, want %q", tt.in, got, tt.want)
		}
		if strings.ContainsAny(got, "+/=") {
			t.Errorf("EncodeToken(%x) = %q contains a character that isn't URL-safe", tt.in, got)
		}
	}
}

func TestDecodeToken(t *testing.T) {
	in := []byte{0xfb, 0xff, 0xfe, 0xff, 0x00, 0x10}
	got, err := DecodeToken(EncodeToken(in))
	if err != nil || !bytes.Equal(got, in) {
		t.Fatalf("round trip gave %x, %v; want %x", got, err, in)
	}
	if _, err := DecodeToken("+//+/w=="); err == nil {
		t.Error("DecodeToken accepted standard (non-URL) base64")
	}
}
