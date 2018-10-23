package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result, index := head, head
	var list []*ListNode
	for nil != index.Next {
		list = append(list, index)
		index = index.Next
	}
	i := len(list) - n
	if i >= 0 {
		list[i].Next = list[i].Next.Next
	} else {
		result = head.Next
		head.Next = nil
	}
	return result
}

func main() {
	l1 := ListNode{6, nil}
	l3 := ListNode{4, &l1}
	l5 := ListNode{2, &l3}

	//l2 := ListNode{3, nil}
	result := removeNthFromEnd(&l5, 3)
	for result != nil {
		println((*result).Val)
		result = (*result).Next
	}
}
