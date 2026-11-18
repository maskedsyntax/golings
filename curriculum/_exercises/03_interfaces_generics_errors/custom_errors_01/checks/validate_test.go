package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidUser(t *testing.T) {
	if err := Validate(User{Name: "Ana", Email: "ana@example.com", Age: 30}); err != nil {
		t.Errorf("Validate(valid user) = %v, want nil", err)
	}
}

func TestInvalidUsers(t *testing.T) {
	for _, tt := range []struct {
		u             User
		field, reason string
	}{
		{User{Email: "a@b", Age: 1}, "name", "must not be empty"},
		{User{Name: "Ana", Email: "ana", Age: 1}, "email", "must contain @"},
		{User{Name: "Ana", Email: "a@b", Age: -1}, "age", "must not be negative"},
		{User{Age: -1}, "name", "must not be empty"}, // the first problem wins
	} {
		err := Validate(tt.u)
		var ve *ValidationError
		if !errors.As(err, &ve) {
			t.Errorf("Validate(%+v) = %v (%T), want a *ValidationError", tt.u, err, err)
			continue
		}
		if ve.Field != tt.field || ve.Reason != tt.reason {
			t.Errorf("Validate(%+v) = %+v, want field %q, reason %q", tt.u, *ve, tt.field, tt.reason)
		}
		if want := tt.field + ": " + tt.reason; err.Error() != want {
			t.Errorf("Error() = %q, want %q", err.Error(), want)
		}
	}
}

func TestAsThroughWrapping(t *testing.T) {
	err := fmt.Errorf("signup: %w", Validate(User{Name: "x", Email: "nope"}))
	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "email" {
		t.Errorf("errors.As through a wrapped error failed: %v", err)
	}
}
