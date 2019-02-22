package main

import "container/list"

//https://leetcode.com/problems/path-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	queue := list.New()
	if nil == root {
		return false
	}
	queue.PushBack(root)
	value := make(map[*TreeNode]int)
	value[root] = root.Val
	for queue.Len() != 0 {
		node := queue.Back().Value.(*TreeNode)
		queue.Remove(queue.Back())
		if nil != node.Left || nil != node.Right {
			if nil != node.Left {
				queue.PushBack(node.Left)
				value[node.Left] = value[node] + node.Left.Val
			}
			if nil != node.Right {
				queue.PushBack(node.Right)
				value[node.Right] = value[node] + node.Right.Val
			}
		} else {
			if value[node] == sum {
				return true
			}
		}
	}
	return false
}

func main() {
	node5 := new(TreeNode)
	node4 := new(TreeNode)
	node8 := new(TreeNode)
	node11 := new(TreeNode)
	node13 := new(TreeNode)
	node4_ := new(TreeNode)
	node7 := new(TreeNode)
	node2 := new(TreeNode)
	node1 := new(TreeNode)
	node5.Val = 5
	node4.Val = 4
	node8.Val = 8
	node11.Val = 11
	node13.Val = 13
	node4_.Val = 4
	node7.Val = 7
	node2.Val = 2
	node1.Val = 1
	node5.Left = node4
	node5.Right = node8
	node4.Left = node11
	node11.Left = node7
	node11.Right = node2
	node8.Left = node13
	node8.Right = node4_
	node4_.Right = node1
	println(hasPathSum(node5, 22))
}
