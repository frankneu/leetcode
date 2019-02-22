package main

import "container/list"

/**
https://leetcode.com/problems/maximum-depth-of-binary-tree/
*/
func maxDepth(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	level := 0
	current := 1
	next := 0
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		current--
		if node.Left != nil {
			queue.PushBack(node.Left)
			next++
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
			next++
		}
		if 0 == current {
			level++
			current = next
			next = 0
		}
	}
	return level
}
