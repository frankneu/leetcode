package main

import "container/list"

/**
基本过程与前序遍历相同，不同之处在于访问的顺序移到出栈时。
1. 找到左子树
2. 左子树为空的时候从栈中取出站定元素，然后遍历
3. 遍历完成指针指向右子树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraverse(root *TreeNode) {
	stack := list.New()
	pNode := root
	for pNode != nil || stack.Len() != 0 {
		if pNode != nil {
			stack.PushBack(pNode)
			pNode = pNode.Left
		} else {
			node := stack.Back().Value.(*TreeNode)
			println(node.Val)
			pNode = node.Right
			stack.Remove(stack.Back())
		}
	}
}
