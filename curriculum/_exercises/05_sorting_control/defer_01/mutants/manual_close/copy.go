// Mutant: closes on every path by hand (correct, but fails the ast check).
package copier

type Resource interface {
	Read() (string, bool)
	Write(s string) error
	Close()
}

func Copy(open, create func() Resource) error {
	src := open()
	dst := create()
	for {
		s, ok := src.Read()
		if !ok {
			break
		}
		if err := dst.Write(s); err != nil {
			dst.Close()
			src.Close()
			return err
		}
	}
	dst.Close()
	src.Close()
	return nil
}
