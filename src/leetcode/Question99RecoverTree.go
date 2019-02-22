package main

/**
https://leetcode.com/problems/recover-binary-search-tree/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Swap(x, y *TreeNode) {
	x.Val, y.Val = y.Val, x.Val
}

//中序遍历，数据存放在管道中
func Inorder(c chan *TreeNode, root *TreeNode) {
	if root == nil {
		return
	}
	Inorder(c, root.Left)
	c <- root
	Inorder(c, root.Right)
}

func recoverTree(root *TreeNode) {
	c := make(chan *TreeNode)
	go func() {
		Inorder(c, root)
		close(c)
	}()

	var first, second *TreeNode
	prev := <-c
	for x := range c {
		if x.Val < prev.Val {
			if first == nil {
				first = prev
			}
			if first != nil {
				second = x
			}
		}
		prev = x
	}
	Swap(first, second)
}
