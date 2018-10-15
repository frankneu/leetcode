package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{6, nil}
	l3 := ListNode{4, &l1}
	l5 := ListNode{2, &l3}

	l2 := ListNode{3, nil}
	result := addTwoNumbers(&l5, &l2)
	for result != nil {
		println((*result).Val)
		result = (*result).Next
	}
}

/**
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	index := new(ListNode)
	index.Val = 0
	index.Next = nil
	var start *ListNode = index
	var before *ListNode = index
	var dig int = 0
	for l1 != nil || l2 != nil {
		var current int = index.Val
		if l1 != nil {
			current += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current += l2.Val
			l2 = l2.Next
		}
		current += dig
		if current > 9 {
			dig = 1
			current -= 10
		} else {
			dig = 0
		}
		index.Val = current
		index.Next = &ListNode{0, nil}
		before = index
		index = index.Next
	}
	if dig == 1 {
		index.Val = dig
		index.Next = nil
		before = index
	}
	if (*index).Val == 0 {
		(*before).Next = nil
	}
	return start
}
