// recover_01: Turning a panic into an error (reference solution)
package safely

import "fmt"

// Safely calls f and converts a panic into an error "recovered: <value>".
func Safely(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r) // err is the named result
		}
	}()
	f()
	return nil
}
