// defer_01: Always clean up (reference solution)
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
	defer src.Close()
	dst := create()
	defer dst.Close()
	for {
		s, ok := src.Read()
		if !ok {
			return nil
		}
		if err := dst.Write(s); err != nil {
			return err // both deferred Close calls still run
		}
	}
}
