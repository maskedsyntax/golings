// recover_02: Recover must be called directly
//
// The panic escapes. Fix Handle. See README.md.
//
// Reference: https://gobyexample.com/recover
package handle

import "fmt"

// capture recovers a panic and stores it in *err.
func capture(err *error) {
	if r := recover(); r != nil {
		*err = fmt.Errorf("handler panicked: %v", r)
	}
}

// Handle runs h and reports a panic as an error.
func Handle(h func() string) (resp string, err error) {
	defer func() {
		capture(&err)
	}()
	return h(), nil
}
