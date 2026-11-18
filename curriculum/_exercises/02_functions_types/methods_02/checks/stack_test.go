package stack

import "testing"

func TestZeroValueIsEmpty(t *testing.T) {
	var s Stack
	if s.Len() != 0 {
		t.Errorf("Len() of a new stack = %d, want 0", s.Len())
	}
	if v, ok := s.Pop(); ok || v != 0 {
		t.Errorf("Pop() on an empty stack = %d, %v, want 0, false", v, ok)
	}
	if v, ok := s.Peek(); ok || v != 0 {
		t.Errorf("Peek() on an empty stack = %d, %v, want 0, false", v, ok)
	}
}

func TestLIFO(t *testing.T) {
	var s Stack
	for _, v := range []int{1, 2, 3} {
		s.Push(v)
	}
	if s.Len() != 3 {
		t.Fatalf("Len() after 3 pushes = %d, want 3", s.Len())
	}
	if v, ok := s.Peek(); !ok || v != 3 {
		t.Errorf("Peek() = %d, %v, want 3, true", v, ok)
	}
	if s.Len() != 3 {
		t.Errorf("Peek changed Len() to %d", s.Len())
	}
	for _, want := range []int{3, 2, 1} {
		if v, ok := s.Pop(); !ok || v != want {
			t.Errorf("Pop() = %d, %v, want %d, true", v, ok, want)
		}
	}
	if s.Len() != 0 {
		t.Errorf("Len() after popping everything = %d, want 0", s.Len())
	}
	if _, ok := s.Pop(); ok {
		t.Error("Pop() on an emptied stack returned ok = true")
	}
}

func TestPushAfterPop(t *testing.T) {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Pop()
	s.Push(9)
	if v, _ := s.Pop(); v != 9 {
		t.Errorf("Pop() = %d, want 9", v)
	}
	if v, _ := s.Pop(); v != 1 {
		t.Errorf("Pop() = %d, want 1", v)
	}
}

func TestZeroIsAValue(t *testing.T) {
	var s Stack
	s.Push(0)
	if v, ok := s.Pop(); !ok || v != 0 {
		t.Errorf("Pop() after Push(0) = %d, %v, want 0, true", v, ok)
	}
}
