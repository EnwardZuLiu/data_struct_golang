package leetcode

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listNode_str1, _ := json.Marshal(l1)
	fmt.Println(string(listNode_str1))
	listNode_str2, _ := json.Marshal(l2)
	fmt.Println(string(listNode_str2))
	var str1 string
	var str2 string
	for i := l1; i != nil; i = i.Next {
		str1 += strconv.Itoa(i.Val)
	}
	for i := l2; i != nil; i = i.Next {
		str2 += strconv.Itoa(i.Val)
	}
	str1 = reverseString(str1)
	str2 = reverseString(str2)
	int1, _ := strconv.ParseInt(str1, 10, 64)
	int2, _ := strconv.ParseInt(str2, 10, 64)

	int3 := int1 + int2

	str3 := strconv.FormatInt(int3, 10)
	str3 = reverseString(str3)

	var topNode *ListNode
	var next *ListNode
	for _, char := range str3 {
		intValue, _ := strconv.Atoi(string(char))
		node := ListNode{Val: intValue, Next: nil}
		if next == nil {
			topNode = &node
			next = &node
			continue
		}
		next.Next = &node
		next = &node
	}
	return topNode
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
