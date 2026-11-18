package pointers

import "testing"

func TestValueOr(t *testing.T) {
	x := 5
	if got := ValueOr(&x, 7); got != 5 {
		t.Errorf("ValueOr(&5, 7) = %d, want 5", got)
	}
	if got := ValueOr(nil, 7); got != 7 {
		t.Errorf("ValueOr(nil, 7) = %d, want 7", got)
	}
	zero := 0
	if got := ValueOr(&zero, 7); got != 0 {
		t.Errorf("ValueOr(&0, 7) = %d, want 0 (a pointer to zero is not nil)", got)
	}
}

func TestIncrement(t *testing.T) {
	x := 41
	Increment(&x)
	if x != 42 {
		t.Errorf("after Increment, x = %d, want 42", x)
	}
	Increment(nil) // must not panic
}

func TestLargest(t *testing.T) {
	a, b, c := 3, 9, -1
	if got := Largest([]*int{&a, nil, &b, &c}); got != &b {
		t.Errorf("Largest returned %v, want the pointer to b", got)
	}
	d, e := 4, 4
	if got := Largest([]*int{&d, &e}); got != &d {
		t.Error("with a tie, Largest should return the first pointer")
	}
	n1, n2 := -5, -2
	if got := Largest([]*int{&n1, &n2}); got != &n2 {
		t.Error("Largest of -5 and -2 should be the pointer to -2")
	}
	if got := Largest(nil); got != nil {
		t.Errorf("Largest(nil) = %v, want nil", got)
	}
	if got := Largest([]*int{nil, nil}); got != nil {
		t.Errorf("Largest([nil nil]) = %v, want nil", got)
	}
}

func TestLargestReturnsCallersPointer(t *testing.T) {
	a, b := 1, 2
	p := Largest([]*int{&a, &b})
	if p == nil {
		t.Fatal("Largest returned nil")
	}
	*p = 100
	if b != 100 {
		t.Error("writing through the returned pointer didn't change b: Largest must return one of the pointers it was given")
	}
}
