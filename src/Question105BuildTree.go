package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func find(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}
	return -1
}

/**
https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 && len(inorder) == 1 {
		node := new(TreeNode)
		node.Val = preorder[0]
		return node
	}
	index := find(inorder, preorder[0])
	root := new(TreeNode)
	root.Val = preorder[0]
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	buildTree(preorder, inorder)
}
