// Mutant: stops at the first problem.
package validate

import "strings"

type User struct {
	Name  string
	Email string
	Age   int
}

type FieldError struct {
	Field  string
	Reason string
}

func (e *FieldError) Error() string { return e.Field + ": " + e.Reason }

func ValidateAll(u User) error {
	switch {
	case u.Name == "":
		return &FieldError{Field: "name", Reason: "must not be empty"}
	case !strings.Contains(u.Email, "@"):
		return &FieldError{Field: "email", Reason: "must contain @"}
	case u.Age < 0:
		return &FieldError{Field: "age", Reason: "must not be negative"}
	}
	return nil
}
