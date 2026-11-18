// Mutant: returns ValidationError values, so errors.As with *ValidationError fails.
package validate

import "strings"

type User struct {
	Name  string
	Email string
	Age   int
}

type ValidationError struct {
	Field  string
	Reason string
}

func (e ValidationError) Error() string { return e.Field + ": " + e.Reason }

func Validate(u User) error {
	switch {
	case u.Name == "":
		return ValidationError{Field: "name", Reason: "must not be empty"}
	case !strings.Contains(u.Email, "@"):
		return ValidationError{Field: "email", Reason: "must contain @"}
	case u.Age < 0:
		return ValidationError{Field: "age", Reason: "must not be negative"}
	}
	return nil
}
