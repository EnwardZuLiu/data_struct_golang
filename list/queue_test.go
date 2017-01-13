package list

import (
	"log"
	"testing"
)

func Test_1(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")
	queue.Enqueue("4")
	queue.Enqueue("5")
	queue.Enqueue("6")
	log.Println(queue.Length())
	log.Println(queue.Dequeue().(string))
	log.Println(queue.Dequeue().(string))
	log.Println(queue.Dequeue().(string))
	log.Println(queue.Dequeue().(string))
	log.Println(queue.Dequeue().(string))
	log.Println(queue.Dequeue().(string))
	log.Println(queue.Length())
}
