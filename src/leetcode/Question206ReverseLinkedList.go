package main

/**
https://leetcode.com/problems/reverse-linked-list/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	previous := head
	head = head.Next
	previous.Next = nil
	for head != nil {
		current := head
		head = head.Next
		current.Next = previous
		previous = current
	}
	return previous
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	next := head.Next
	if next != nil {
		ans := reverseListRecursive(next)
		next.Next = head
		head.Next = nil
		return ans
	} else {
		head.Next = nil
		return head
	}
}

func build(nums []int) *ListNode {
	var head *ListNode
	var pre *ListNode
	var next *ListNode
	for i := 0; i < len(nums); i++ {
		next = new(ListNode)
		next.Val = nums[i]
		if i == 0 {
			head = next
		} else {
			pre.Next = next
		}
		pre = next
	}
	return head
}

func main() {
	head := build([]int{1, 2})
	head = reverseListRecursive(head)
	println(head.Val)
}
