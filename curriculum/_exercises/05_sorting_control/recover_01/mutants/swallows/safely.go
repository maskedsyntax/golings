// Mutant: recovers but reports no error.
package safely

func Safely(f func()) error {
	defer func() { recover() }()
	f()
	return nil
}
