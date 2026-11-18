package container

import "testing"

func TestPromotion(t *testing.T) {
	c := Container{Base: Base{ID: 7}, Items: []string{"a", "b"}}
	if c.ID != 7 {
		t.Errorf("c.ID = %d, want 7", c.ID)
	}
	c.ID = 8
	if c.Base.ID != 8 {
		t.Errorf("after c.ID = 8, c.Base.ID = %d: the promoted field is the embedded field", c.Base.ID)
	}
	if got := c.Describe(); got != "base with id=8" {
		t.Errorf("c.Describe() = %q, want %q", got, "base with id=8")
	}
}

func TestSatisfiesInterface(t *testing.T) {
	var d Describer = Container{Base: Base{ID: 1}}
	if got := d.Describe(); got != "base with id=1" {
		t.Errorf("Describe() via the interface = %q", got)
	}
}
