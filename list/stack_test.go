package list

import (
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
	if stack.Length() == 7 && stack.top.data == "liuzm6" {
		t.Log("success")
	} else {
		t.Error("fail")
	}
	for stack.Length() > 0 {
		t.Log(stack.Pop().(string))
	}
	if stack.IsEmpty() {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
