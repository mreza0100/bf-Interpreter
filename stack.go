package main

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	l := len(*s)
	if l > 0 {
		op := (*s)[l-1]
		*s = (*s)[:l-1]
		return op
	}
	return ""
}

func (s *Stack) Top() string {
	n := len(*s) - 1
	return (*s)[n]
}

func (s Stack) Len() int {
	return len(s)
}
