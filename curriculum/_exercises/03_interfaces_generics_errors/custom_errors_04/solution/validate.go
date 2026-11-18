// custom_errors_04: Reporting every problem (reference solution)
package validate

import (
	"errors"
	"strings"
)

// User is a sign-up form.
type User struct {
	Name  string
	Email string
	Age   int
}

// FieldError reports one invalid field.
type FieldError struct {
	Field  string
	Reason string
}

func (e *FieldError) Error() string { return e.Field + ": " + e.Reason }

// ValidateAll returns every problem with u joined into one error, or nil.
func ValidateAll(u User) error {
	var errs []error
	if u.Name == "" {
		errs = append(errs, &FieldError{Field: "name", Reason: "must not be empty"})
	}
	if !strings.Contains(u.Email, "@") {
		errs = append(errs, &FieldError{Field: "email", Reason: "must contain @"})
	}
	if u.Age < 0 {
		errs = append(errs, &FieldError{Field: "age", Reason: "must not be negative"})
	}
	return errors.Join(errs...) // nil when errs is empty
}
