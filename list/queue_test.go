package list

import (
	"testing"
)

func Test_1(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue("1")
	if queue.Length() == 1 && queue.data.data == "1" {
		t.Log("测试通过")
	} else {
		t.Error("Push 函数测试失败")
	}
	queue.Enqueue("2")
	queue.Enqueue("3")
	queue.Enqueue("4")
	queue.Enqueue("5")
	queue.Enqueue("6")
	if queue.Length() == 6 && queue.data.data == "1" {
		t.Log("success")
	} else {
		t.Error("fail")
	}
	if queue.Dequeue() == "1" {
		t.Log("success")
	} else {
		t.Error("fail")
	}
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	if queue.IsEmpty() {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
