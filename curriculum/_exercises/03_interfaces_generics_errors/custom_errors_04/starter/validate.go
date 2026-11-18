// custom_errors_04: Reporting every problem
//
// Implement ValidateAll with errors.Join. See README.md.
//
//	ValidateAll(User{Email: "x", Age: -1}).Error() ==
//		"name: must not be empty\nemail: must contain @\nage: must not be negative"
//
// Reference: https://gobyexample.com/custom-errors
package validate

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
	// TODO
	return nil
}
