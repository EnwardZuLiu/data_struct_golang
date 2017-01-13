package list

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	data interface{}
	next *Node
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(value interface{}) {
	s.top = &Node{value, s.top}
	s.size++
}

func (s *Stack) Peek() interface{} {
	return s.top.data
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		val := s.top.data
		s.top = s.top.next
		s.size--
		return val
	}
	return nil
}
