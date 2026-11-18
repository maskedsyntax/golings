// methods_02: A stack with a useful zero value (reference solution)
package stack

// Stack is a LIFO stack of ints. The zero value is an empty stack.
type Stack struct {
	items []int
}

// Push adds v to the top of the stack.
func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

// Pop removes and returns the top value. ok is false if the stack is empty.
func (s *Stack) Pop() (v int, ok bool) {
	v, ok = s.Peek()
	if ok {
		s.items = s.items[:len(s.items)-1]
	}
	return v, ok
}

// Peek returns the top value without removing it. ok is false if the stack is empty.
func (s *Stack) Peek() (v int, ok bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// Len returns the number of values on the stack.
func (s *Stack) Len() int {
	return len(s.items)
}
