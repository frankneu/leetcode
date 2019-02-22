package main

import "container/list"

/**
https://leetcode.com/problems/symmetric-tree/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	left := list.New()
	right := list.New()
	if nil != root.Left {
		left.PushBack(root.Left)
	}
	if nil != root.Right {
		right.PushBack(root.Right)
	}
	for left.Len() != 0 && right.Len() != 0 {
		l := left.Front().Value.(*TreeNode)
		r := right.Front().Value.(*TreeNode)
		left.Remove(left.Front())
		right.Remove(right.Front())
		if l.Val != r.Val {
			return false
		}
		if l.Left != nil && r.Right != nil {
			left.PushBack(l.Left)
			right.PushBack(r.Right)
		} else if l.Left != nil || r.Right != nil {
			return false
		}
		if l.Right != nil && r.Left != nil {
			left.PushBack(l.Right)
			right.PushBack(r.Left)
		} else if l.Right != nil || r.Left != nil {
			return false
		}

	}
	return left.Len() == 0 && right.Len() == 0
}
