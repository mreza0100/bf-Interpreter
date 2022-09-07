package brainfuck

type loopStack struct {
	data []int
}

func newLoopStack() *loopStack {
	return &loopStack{
		data: make([]int, 0),
	}
}

func (s *loopStack) push(i int) {
	s.data = append(s.data, i)
}

func (s *loopStack) pop() int {
	if s.len() == 0 {
		panic("stack is empty")
	}

	i := s.data[s.len()-1]
	s.data = s.data[:s.len()-1]
	return i
}

func (s *loopStack) isEmpty() bool {
	return s.len() == 0
}

func (s *loopStack) len() int {
	return len(s.data)
}
