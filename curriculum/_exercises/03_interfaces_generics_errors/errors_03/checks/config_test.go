package config

import (
	"errors"
	"testing"
)

func TestFound(t *testing.T) {
	if v, err := GetOrDefault("name", "x"); err != nil || v != "golings" {
		t.Errorf("GetOrDefault(\"name\") = %q, %v, want \"golings\", nil", v, err)
	}
}

func TestMissingGivesDefault(t *testing.T) {
	if v, err := GetOrDefault("missing", "fallback"); err != nil || v != "fallback" {
		t.Errorf("GetOrDefault(\"missing\", \"fallback\") = %q, %v, want \"fallback\", nil", v, err)
	}
}

func TestOtherErrorsAreReturned(t *testing.T) {
	v, err := GetOrDefault("locked", "fallback")
	if !errors.Is(err, ErrPermission) {
		t.Errorf("GetOrDefault(\"locked\") error = %v, want an error wrapping ErrPermission", err)
	}
	if v != "" {
		t.Errorf("GetOrDefault(\"locked\") returned %q alongside the error, want \"\"", v)
	}
}
