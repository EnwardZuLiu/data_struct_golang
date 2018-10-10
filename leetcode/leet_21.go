package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	var l3 *ListNode
	var point *ListNode
	for l1 != nil || l2 != nil {
		if l1 == nil && l2 != nil {
			point.Next = l2
			return l3
		}
		if l1 != nil && l2 == nil {
			point.Next = l1
			return l3
		}
		if l1.Val <= l2.Val {
			node := &ListNode{Val: l1.Val, Next: nil}
			if l3 == nil {
				l3 = node
				point = node
			}
			point.Next = node
			point = node
			l1 = l1.Next
		} else {
			node := &ListNode{Val: l2.Val, Next: nil}
			if l3 == nil {
				l3 = node
				point = node
			}
			point.Next = node
			point = node
			l2 = l2.Next
		}
	}
	return l3
}
