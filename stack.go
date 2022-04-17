package brainfuck

// LoopStack for loops
type LoopStack struct {
	data []int
}

func NewLoopStack() *LoopStack {
	return &LoopStack{
		data: make([]int, 0),
	}
}

func (s *LoopStack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *LoopStack) Pop() int {
	if len(s.data) == 0 {
		panic("stack is empty")
	}

	i := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return i
}

func (s *LoopStack) Top() int {
	if len(s.data) == 0 {
		panic("stack is empty")
	}

	return s.data[len(s.data)-1]
}

func (s *LoopStack) IsEmpty() bool {
	return len(s.data) == 0
}
