// Mutant: joins the messages into a plain string error, losing the FieldErrors.
package validate

import (
	"errors"
	"strings"
)

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
	var msgs []string
	if u.Name == "" {
		msgs = append(msgs, "name: must not be empty")
	}
	if !strings.Contains(u.Email, "@") {
		msgs = append(msgs, "email: must contain @")
	}
	if u.Age < 0 {
		msgs = append(msgs, "age: must not be negative")
	}
	if len(msgs) == 0 {
		return nil
	}
	return errors.New(strings.Join(msgs, "\n"))
}
