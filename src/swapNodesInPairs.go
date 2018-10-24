package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var result *ListNode
	if nil != head && nil != head.Next {
		result = head.Next
	} else {
		return head
	}
	var pre *ListNode
	index := head
	for nil != index {
		a := index
		b := index.Next
		if nil == b {
			return result
		}
		if nil != pre {
			pre.Next = b
		}
		index = b.Next
		b.Next = a
		a.Next = index
		pre = a
	}
	return result
}

func main() {
	l1 := ListNode{5, nil}
	l2 := ListNode{4, &l1}
	l3 := ListNode{3, &l2}
	l4 := ListNode{2, &l3}
	l5 := ListNode{1, &l4}
	result := &l5
	//l2 := ListNode{3, nil}
	result = swapPairs(&l5)
	for result != nil {
		println((*result).Val)
		result = (*result).Next
	}
}
