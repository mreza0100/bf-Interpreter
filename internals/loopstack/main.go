package loopstack

type loopstack struct {
	data []int
}

type Loopstack interface {
	Push(i int)
	Pop() int
	IsEmpty() bool
	Len() int
}

func New() Loopstack {
	return &loopstack{data: make([]int, 0, 3)}
}

func (s *loopstack) Push(i int) {
	s.data = append(s.data, i)
}

func (s *loopstack) Pop() int {
	if s.Len() == 0 {
		panic("stack is empty")
	}

	i := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return i
}

func (s *loopstack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *loopstack) Len() int {
	return len(s.data)
}
