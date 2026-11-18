// custom_errors_01: An error type with fields
//
// Implement Error and Validate. See README.md.
//
//	err := Validate(User{Name: "Ana", Email: "ana"})
//	err.Error() == "email: must contain @"
//
// Reference: https://gobyexample.com/custom-errors
package validate

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

// TODO: func (e *ValidationError) Error() string

// Validate returns a *ValidationError for the first invalid field, or nil.
func Validate(u User) error {
	// TODO
	return nil
}
