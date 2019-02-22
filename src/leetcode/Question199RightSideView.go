package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/binary-tree-right-side-view/
*/

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if nil == root {
		return result
	}
	queue := list.New()
	next := list.New()
	queue.PushBack(root)
	nums := make([]int, 0)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		nums = append(nums, node.Val)
		if nil != node.Left {
			next.PushBack(node.Left)
		}
		if nil != node.Right {
			next.PushBack(node.Right)
		}
		if 0 == queue.Len() {
			result = append(result, nums[len(nums)-1])
			queue = next
			next = list.New()
		}
	}
	return result
}
