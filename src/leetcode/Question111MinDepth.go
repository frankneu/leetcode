package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/minimum-depth-of-binary-tree/
*/
func minDepth(root *TreeNode) int {
	queue := list.New()
	next := list.New()
	if root == nil {
		return 0
	}
	level := 1
	queue.PushBack(root)
	for queue.Len() >= 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if node.Left == nil && node.Right == nil {
			return level
		}
		if node.Left != nil {
			next.PushBack(node.Left)
		}
		if node.Right != nil {
			next.PushBack(node.Right)
		}
		if queue.Len() == 0 {
			level++
			queue = next
			next = list.New()
		}
	}
	return -1
}
