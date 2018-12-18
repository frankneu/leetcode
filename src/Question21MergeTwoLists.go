package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var index *ListNode
	for l1 != nil && l2 != nil {
		var node *ListNode
		if l1.Val > l2.Val {
			node = l2
			l2 = l2.Next
		} else {
			node = l1
			l1 = l1.Next
		}
		if head == nil {
			head = node
			index = node
		} else {
			index.Next = node
			index = node
		}
	}
	var remain *ListNode
	if l2 != nil {
		remain = l2
	}
	if l1 != nil {
		remain = l1
	}
	if head == nil {
		head = remain
	} else {
		index.Next = remain
	}
	return head
}
