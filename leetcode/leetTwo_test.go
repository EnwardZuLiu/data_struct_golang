package leetcode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	node1 := ListNode{Val: 5}
	node2 := ListNode{Val: 6}
	node3 := ListNode{Val: 7}
	node4 := ListNode{Val: 2}
	node5 := ListNode{Val: 4}
	node6 := ListNode{Val: 6}
	node7 := ListNode{Val: 7}
	node8 := ListNode{Val: 9}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	node5.Next = &node6
	node6.Next = &node7
	node7.Next = &node8

	l := addTwoNumbers(&node1, &node5)

	s, _ := json.Marshal(l)

	fmt.Println(string(s))
}
