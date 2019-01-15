package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/find-bottom-left-tree-value/
*/
func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	count := 1
	next := 0
	result := root.Val
	flag := true
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if flag {
			result = node.Val
			flag = false
		}
		count--
		if node.Left != nil {
			queue.PushBack(node.Left)
			next++
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
			next++
		}
		if count == 0 {
			count = next
			flag = true
			next = 0
		}
	}
	return result
}
