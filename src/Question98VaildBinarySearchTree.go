package main

import (
	"container/list"
	"math"
)

/**
https://leetcode.com/problems/validate-binary-search-tree/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root != root {
		return true
	}
	queue := list.New()
	node := root
	source := math.MinInt64
	for node != nil || queue.Len() != 0 {
		for nil != node {
			queue.PushFront(node)
			node = node.Left
		}
		if queue.Len() != 0 {
			node = queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			if node.Val <= source {
				return false
			}
			source = node.Val
			node = node.Right
		}
	}
	return true
}

func main() {
	//node5 := new(TreeNode)
	//node4 := new(TreeNode)
	//node6 := new(TreeNode)
	//node1 := new(TreeNode)
	//node3 := new(TreeNode)
	//node5.Val = 5
	//node4.Val = 4
	//node6.Val = 6
	//node1.Val = 1
	//node3.Val = 3
	//node5.Left = node1
	//node5.Right = node4
	//node4.Left = node3
	//node4.Right = node6
	//node2 := new(TreeNode)
	//node1 := new(TreeNode)
	//node3 := new(TreeNode)
	//node2.Val = 2
	//node1.Val = 1
	//node3.Val = 3
	//node2.Left = node1
	//node2.Right = node3
	node0 := new(TreeNode)
	node1 := new(TreeNode)
	node2 := new(TreeNode)
	node3 := new(TreeNode)
	node4 := new(TreeNode)
	node5 := new(TreeNode)
	node6 := new(TreeNode)
	node0.Val = 0
	node1.Val = 1
	node2.Val = 2
	node3.Val = 3
	node4.Val = 4
	node5.Val = 5
	node6.Val = 6
	node3.Left = node1
	node3.Right = node5
	node1.Left = node0
	node1.Right = node2
	node5.Left = node4
	node5.Right = node6
	node := new(TreeNode)
	node.Val = 3
	node2.Right = node
	println(isValidBST(node3))
}
