// defer_01: Always clean up
//
// Copy leaks resources when a write fails. Use defer. See README.md.
//
// Reference: https://gobyexample.com/defer
package copier

// Resource is something that must be closed.
type Resource interface {
	Read() (string, bool)
	Write(s string) error
	Close()
}

// Copy copies every value from the source to the destination.
func Copy(open, create func() Resource) error {
	src := open()
	dst := create()
	for {
		s, ok := src.Read()
		if !ok {
			break
		}
		if err := dst.Write(s); err != nil {
			return err
		}
	}
	dst.Close()
	src.Close()
	return nil
}
