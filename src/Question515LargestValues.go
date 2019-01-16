package main

import (
	"container/list"
	"math"
)

/**
https://leetcode.com/problems/find-largest-value-in-each-tree-row/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	result := make([]int, 0)
	if nil == root {
		return result
	}
	queue := list.New()
	queue.PushBack(root)
	row := 1
	next := 0
	max := math.MinInt32
	for queue.Len() != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if max < node.Val {
			max = node.Val
		}
		row--
		if nil != node.Left {
			queue.PushBack(node.Left)
			next++
		}
		if nil != node.Right {
			queue.PushBack(node.Right)
			next++
		}
		if row == 0 {
			row = next
			next = 0
			result = append(result, max)
			max = math.MinInt32
		}
	}
	return result
}
