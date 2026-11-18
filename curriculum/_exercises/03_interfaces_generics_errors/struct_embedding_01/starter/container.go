// struct_embedding_01: Promoted fields and methods
//
// The tests don't compile: embed Base in Container.
//
//	c := Container{Base: Base{ID: 7}, Items: []string{"a"}}
//	c.ID == 7; c.Describe() == "base with id=7"
//
// Reference: https://gobyexample.com/struct-embedding
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
	base  Base
	Items []string
}
