package list

type Queue struct {
	size int
	data *Node
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) Enqueue(d interface{}) {
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

func (q *Queue) Dequeue() interface{} {
	if q.Length() == 0 {
		return nil
	}
	val := q.data.data
	q.data = q.data.next
	q.size--
	return val
}

/**
 * 对队列中的每一个元素进行 visit 操作
 */
func (q *Queue) Traverse(visit func(interface{})) {

}
