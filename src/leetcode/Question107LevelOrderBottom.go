package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
*/
func levelOrderBottom(root *TreeNode) [][]int {
	queue := list.New()
	if root == nil {
		return [][]int{}
	}
	queue.PushBack(root)
	next := list.New()
	value := make([]int, 0)
	var result [][]int
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		value = append(value, node.Val)
		if nil != node.Left {
			next.PushBack(node.Left)
		}
		if nil != node.Right {
			next.PushBack(node.Right)
		}
		if queue.Len() == 0 {
			queue = next
			next = list.New()
			result = append([][]int{value}, result...)
			value = make([]int, 0)
		}
	}
	return result
}

func main() {
	node3 := new(TreeNode)
	node9 := new(TreeNode)
	node20 := new(TreeNode)
	node15 := new(TreeNode)
	node7 := new(TreeNode)
	node3.Val = 3
	node9.Val = 9
	node20.Val = 20
	node15.Val = 15
	node7.Val = 7
	node3.Left = node9
	node3.Right = node20
	node20.Left = node15
	node20.Right = node7
	levelOrderBottom(node3)
}
