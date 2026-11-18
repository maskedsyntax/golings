package handle

import (
	"strings"
	"testing"
)

func TestHandleOK(t *testing.T) {
	resp, err := Handle(func() string { return "200 OK" })
	if resp != "200 OK" || err != nil {
		t.Fatalf("Handle = %q, %v; want \"200 OK\", nil", resp, err)
	}
}

func TestHandleRecovers(t *testing.T) {
	var resp string
	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("the panic escaped from Handle: %v", r)
			}
		}()
		resp, err = Handle(func() string { panic("nil user") })
	}()
	if err == nil || !strings.Contains(err.Error(), "nil user") || resp != "" {
		t.Fatalf("Handle(panicking) = %q, %v; want \"\" and an error mentioning the panic", resp, err)
	}
}
