package list

/**
 * 双端队列的实现
 * 可以从队头或者队尾任意进出
 */

// deque
type Deque struct {
	size int
	top  *Node
	end  *Node
}

// length
func (d *Deque) length() int {
	return d.size
}

// push of top
func (d *Deque) PushTop(value interface{}) interface{} {

	if value == nil {
		return nil
	}

	node := &Node{value, nil}
	if d.size == 0 {
		d.top = node
		d.end = node
		d.size = 1
		return value
	}

	node.next = d.top
	d.top = node
	d.size = d.size + 1
	return value
}

// push of end
func (d *Deque) PushEnd(value interface{}) interface{} {
	if value == nil {
		return nil
	}

	node := &Node{value, nil}
	if d.size == 0 {
		d.top = node
		d.end = node
		d.size = 1
		return value
	}

	d.end.next = node
	d.end = node
	return node
}

func (d *Deque) PopTop() interface{} {
	if d.size == 0 {
		return nil
	}

	value := d.top.data
	d.top = d.top.next
	d.size = d.size - 1
	return value
}

func (d *Deque) PopEnd() interface{} {
	if d.size == 0 {
		return nil
	}
	// TODO
	return nil
}
