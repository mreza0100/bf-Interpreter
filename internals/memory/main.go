package memory

type memory struct {
	values   []byte
	isStatic bool
}

type Memory interface {
	Len() int
	IsOut(idx int) bool
	IncreaseCap()
	GetMemory() []byte
	IsStatic() bool
}

func New(initialSize int, isStatic bool) Memory {
	return &memory{
		values:   make([]byte, initialSize),
		isStatic: isStatic,
	}
}

func (m *memory) IsStatic() bool {
	return m.isStatic
}

func (m *memory) GetMemory() []byte {
	return m.values
}

func (m *memory) Len() int {
	return len(m.values)
}

func (m *memory) IsOut(idx int) bool {
	return idx < 0 || idx >= len(m.values)
}

func (m *memory) IncreaseCap() {
	if m.isStatic {
		return
	}

	// 2 * m.len()
	m.values = append(m.values, make([]byte, len(m.values))...)
}
