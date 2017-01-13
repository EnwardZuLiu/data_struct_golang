package main

import (
	"fmt"
)

type Queue struct {
	size int
	data *Node
}

type Node struct {
	data interface{}
	next *Node
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) Push(d interface{}) {
	n := &Node{d, nil}
	if q.Length() == 0 {
		q.data = n
		q.size++
		return
	}
	l := q.data
	length := q.Length()
	for i := 0; i < length; i++ {
		if i == q.Length()-1 {
			l.next = n
			q.size++
		} else {
			l = l.next
		}
	}
}

func (q *Queue) Pop() interface{} {
	if q.Length() == 0 {
		return nil
	}
	val := q.data.data
	q.data = q.data.next
	q.size--
	return val
}

func main() {
	queue := new(Queue)
	queue.Push("1")
	queue.Push("2")
	queue.Push("3")
	queue.Push("4")
	queue.Push("5")
	queue.Push("6")
	fmt.Println(queue.Length())
	fmt.Println(queue.Pop().(string))
	fmt.Println(queue.Pop().(string))
	fmt.Println(queue.Pop().(string))
	fmt.Println(queue.Pop().(string))
	fmt.Println(queue.Pop().(string))
	fmt.Println(queue.Pop().(string))
	fmt.Println(queue.Length())
}
