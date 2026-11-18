// struct_embedding_01: Promoted fields and methods (reference solution)
package container

import "fmt"

// Describer can describe itself.
type Describer interface {
	Describe() string
}

// Base holds what every stored object has.
type Base struct {
	ID int
}

// Describe returns a short description of b.
func (b Base) Describe() string {
	return fmt.Sprintf("base with id=%d", b.ID)
}

// Container is a Base that also holds items.
type Container struct {
	Base
	Items []string
}
