// methods_02: A stack with a useful zero value
//
// Implement the Stack methods. See README.md.
//
//	var s Stack
//	s.Push(1); s.Push(2)
//	s.Pop() == 2, true
//	s.Len() == 1
//
// Reference: https://gobyexample.com/methods
package stack

// Stack is a LIFO stack of ints. The zero value is an empty stack.
type Stack struct {
	items []int
}

// Push adds v to the top of the stack.
func (s *Stack) Push(v int) {
	// TODO
}

// Pop removes and returns the top value. ok is false if the stack is empty.
func (s *Stack) Pop() (v int, ok bool) {
	// TODO
	return 0, false
}

// Peek returns the top value without removing it. ok is false if the stack is empty.
func (s *Stack) Peek() (v int, ok bool) {
	// TODO
	return 0, false
}

// Len returns the number of values on the stack.
func (s *Stack) Len() int {
	// TODO
	return 0
}
