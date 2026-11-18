// recover_01: Turning a panic into an error
//
// Implement Safely. See README.md.
//
// Reference: https://gobyexample.com/recover
package safely

// Safely calls f and converts a panic into an error "recovered: <value>".
func Safely(f func()) error {
	// TODO
	f()
	return nil
}
