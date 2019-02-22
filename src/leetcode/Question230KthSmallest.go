package main

import "container/list"

/**
https://leetcode.com/problems/kth-smallest-element-in-a-bst/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	ans := make([]int, 0)
	if root.Left != nil {
		ans = append(ans, inorder(root.Left)...)
	}
	ans = append(ans, root.Val)
	if root.Right != nil {
		ans = append(ans, inorder(root.Right)...)
	}
	return ans
}

func kthSmallest(root *TreeNode, k int) int {
	return inorder(root)[k-1]
}
