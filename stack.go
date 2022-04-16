package main

type stack struct {
	next     *stack
	startIdx int
	endIdx   int
}

func (s *stack) getNext() *stack {
	return s.next
}

func (s *stack) push(newStack stack) {
	newStack.next = s
	s = &newStack
}

func (s *stack) destory() {
	for s != nil {
		s = s.next
	}
}

// stack
// ^
// |
// stack
// ^
// |
// stack
