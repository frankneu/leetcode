package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
参考了提示，这个结果实际上就是前序遍历的结果
*/
func preOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	result := []*TreeNode{root}
	if root.Left != nil {
		value := preOrder(root.Left)
		result = append(result, value...)
	}
	if root.Right != nil {
		value := preOrder(root.Right)
		result = append(result, value...)
	}
	return result
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	order := preOrder(root)
	for i := 1; i < len(order); i++ {
		order[i-1].Left = nil
		order[i-1].Right = order[i]
	}
}
