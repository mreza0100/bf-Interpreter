package brainfuck

type memory struct {
	values   []byte
	isStatic bool
}

func newMemory(initialSize int, isStatic bool) *memory {
	m := new(memory)

	m.values = make([]byte, initialSize)
	m.isStatic = isStatic

	return m
}

func (m *memory) len() int {
	return len(m.values)
}

func (m *memory) isOut(idx int) bool {
	return idx < 0 || idx >= len(m.values)
}

func (m *memory) moreCap() {
	if m.isStatic {
		return
	}

	// 2 * m.len()
	m.values = append(m.values, make([]byte, len(m.values))...)
}
