package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func balanceDepth(root *TreeNode) (bool, int) {
	if nil == root.Left && nil == root.Right {
		return true, 1
	}
	balLeft := true
	depthLeft := 0
	balRight := true
	depthRight := 0
	if nil != root.Left {
		balLeft, depthLeft = balanceDepth(root.Left)
	}
	if nil != root.Right {
		balRight, depthRight = balanceDepth(root.Right)
	}
	if balLeft && balRight {
		if math.Abs(float64(depthLeft-depthRight)) > 1 {
			return false, 0
		} else {
			return true, max(depthLeft, depthRight) + 1
		}
	} else {
		return false, 0
	}
}

/**
https://leetcode.com/problems/balanced-binary-tree/
*/
func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	balance, _ := balanceDepth(root)
	return balance
}
