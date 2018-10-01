package list

import "errors"

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	data interface{}
	next *Node
}

/**
 * destory stack , stack is nil
 */
func (s *Stack) DestoryStack() {
	if s != nil {
		s = nil
	}
}

/**
 * clear stack, stack node is nil and stack length is zero
 */
func (s *Stack) ClearStack() {
	if s != nil {
		s.top = nil
		s.size = 0
	}
}

/**
 * stack length
 */
func (s *Stack) Length() (int, error) {
	if s == nil {
		return 0, errors.New("stack is not exist")
	}
	return s.size, nil
}

/**
 * return top element, bug do not delete top element
 */
func (s *Stack) GetTop() (interface{}, error) {
	if s == nil {
		return nil, errors.New("stack is no exist")
	}
	return s.top.data, nil
}

/**
 *
 */
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

/**
 * into stack
 */
func (s *Stack) Push(value interface{}) {
	s.top = &Node{value, s.top}
	s.size++
}

func (s *Stack) Peek() interface{} {
	return s.top.data
}

/**
 * out stack
 */
func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		val := s.top.data
		s.top = s.top.next
		s.size--
		return val
	}
	return nil
}

/**
 * 对 stack 中的每一个元素执行某个操作
 */
func (s *Stack) Traverse(visit func(interface{})) error {
	if s == nil {
		return errors.New("stack is not exist")
	}
	node := s.top

	for node != nil {
		visit(node.data)
		node = node.next
	}
	return nil
}
