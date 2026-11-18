// Mutant: pops from the bottom (a queue, not a stack).
package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) { s.items = append(s.items, v) }

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	v := s.items[0]
	s.items = s.items[1:]
	return v, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[0], true
}

func (s *Stack) Len() int { return len(s.items) }
