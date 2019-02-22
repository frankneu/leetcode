package main

/**
https://leetcode.com/problems/binary-tree-maximum-path-sum/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if 0 == len(nums) {
		return nil
	}
	node := new(TreeNode)
	m := (len(nums) - 1) / 2
	node.Val = nums[m]
	if m > 0 {
		node.Left = sortedArrayToBST(nums[:m])
	}
	if m < len(nums)-1 {
		node.Right = sortedArrayToBST(nums[m:])
	}
	return node
}
