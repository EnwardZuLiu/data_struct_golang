package list

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	stack := new(Stack)
	stack.Push("liuzm")
	stack.Push("liuzm1")
	stack.Push("liuzm2")
	stack.Push("liuzm3")
	stack.Push("liuzm4")
	stack.Push("liuzm5")
	stack.Push("liuzm6")
	log.Println(stack.Length())
	for stack.Length() > 0 {
		log.Printf("%s \n", stack.Pop().(string))
		log.Println(stack.Length())
	}
	log.Println()
}
