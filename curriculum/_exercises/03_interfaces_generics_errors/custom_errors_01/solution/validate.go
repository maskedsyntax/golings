// custom_errors_01: An error type with fields (reference solution)
package validate

import "strings"

// User is a sign-up form.
type User struct {
	Name  string
	Email string
	Age   int
}

// ValidationError reports which field is invalid and why.
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Reason
}

// Validate returns a *ValidationError for the first invalid field, or nil.
func Validate(u User) error {
	switch {
	case u.Name == "":
		return &ValidationError{Field: "name", Reason: "must not be empty"}
	case !strings.Contains(u.Email, "@"):
		return &ValidationError{Field: "email", Reason: "must contain @"}
	case u.Age < 0:
		return &ValidationError{Field: "age", Reason: "must not be negative"}
	}
	return nil
}
