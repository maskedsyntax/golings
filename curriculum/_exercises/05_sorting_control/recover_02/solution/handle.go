// recover_02: Recover must be called directly (reference solution)
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
	defer capture(&err) // capture is the deferred function, so its recover works
	return h(), nil
}
