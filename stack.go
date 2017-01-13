package main

import "fmt"

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

func main() {
	stack := new(Stack)
	stack.Push("liuzm")
	stack.Push("liuzm1")
	stack.Push("liuzm2")
	stack.Push("liuzm3")
	stack.Push("liuzm4")
	stack.Push("liuzm5")
	stack.Push("liuzm6")
	fmt.Println(stack.Length())
	for stack.Length() > 0 {
		fmt.Printf("%s \n", stack.Pop().(string))
		fmt.Println(stack.Length())
	}
	fmt.Println()
}
