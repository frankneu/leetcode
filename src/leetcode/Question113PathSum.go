package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/path-sum-ii/
*/

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	if nil == root {
		return result
	}
	stack := list.New()
	stack.PushBack(root)
	value := make(map[*TreeNode][]int)
	value[root] = []int{root.Val}
	for stack.Len() != 0 {
		node := stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		if nil != node.Left {
			left := make([]int, 0)
			left = append(left, value[node]...)
			value[node.Left] = append(left, node.Left.Val)
			stack.PushBack(node.Left)
		}
		if nil != node.Right {
			right := make([]int, 0)
			right = append(right, value[node]...)
			value[node.Right] = append(right, node.Right.Val)
			stack.PushBack(node.Right)
		}
		if nil == node.Left && nil == node.Right {
			target := 0
			for i := 0; i < len(value[node]); i++ {
				target += value[node][i]
			}
			if sum == target {
				result = append(result, value[node])
			}
		}
	}
	return result
}

func main() {
	pathSum(nil, 1)
}
