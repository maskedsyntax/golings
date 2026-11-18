// Mutant: recovers into a copy of err, so the caller never sees it.
package handle

import "fmt"

func capture(err error) {
	if r := recover(); r != nil {
		err = fmt.Errorf("handler panicked: %v", r)
		_ = err
	}
}

func Handle(h func() string) (resp string, err error) {
	defer capture(err)
	return h(), nil
}
