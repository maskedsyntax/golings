// Mutant: Container defines its own Describe, hiding Base's.
package container

import "fmt"

type Describer interface {
	Describe() string
}

type Base struct {
	ID int
}

func (b Base) Describe() string { return fmt.Sprintf("base with id=%d", b.ID) }

type Container struct {
	Base
	Items []string
}

func (c Container) Describe() string { return fmt.Sprintf("container with %d items", len(c.Items)) }
