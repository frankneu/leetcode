package main

import "container/list"

/**
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var queue = list.New()
	if nil == root {
		return [][]int{}
	}
	queue.PushFront(root)
	var nextQueue = list.New()
	var result [][]int
	var value []int = make([]int, 0)
	flag := true
	for queue.Len() > 0 {
		element := (queue.Front().Value).(*TreeNode)
		queue.Remove(queue.Front())
		if flag {
			value = append(value, element.Val)
		} else {
			value = append([]int{element.Val}, value...)
		}
		if nil != element.Left {
			nextQueue.PushBack(element.Left)
		}
		if nil != element.Right {
			nextQueue.PushBack(element.Right)
		}
		if queue.Len() == 0 {
			queue = nextQueue
			nextQueue = list.New()
			flag = !flag
			result = append(result, value)
			value = make([]int, 0)
		}
	}
	return result
}
