package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
*/

func find(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}
	return -1
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		node := new(TreeNode)
		node.Val = inorder[0]
		return node
	}
	root := postorder[len(postorder)-1]
	node := new(TreeNode)
	index := find(inorder, root)
	node.Val = root
	node.Left = buildTree(inorder[:index], postorder[:index])
	node.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return node
}
