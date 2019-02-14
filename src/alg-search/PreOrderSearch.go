package main

import "container/list"

/**
根据前序遍历的顺序，优先访问根结点，然后在访问左子树和右子树。
所以，对于任意结点node，第一部分即直接访问之，之后在判断左子树是否为空，不为空时即重复上面的步骤，直到其为空。
若为空，则需要访问右子树。
注意，在访问过左孩子之后，需要反过来访问其右孩子，所以，需要栈这种数据结构的支持。对于任意一个结点node，具体步骤如下：
a)访问之，并把结点node入栈，当前结点置为左孩子；
b)判断结点node是否为空，
	若为空，则取出栈顶结点并出栈，将右孩子置为当前结点；
	否则重复a)步直到当前结点为空或者栈为空（可以发现栈中的结点就是为了访问右孩子才存储的）
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraverse(root *TreeNode) {
	stack := list.New()
	pNode := root
	for pNode != nil || stack.Len() != 0 {
		if pNode != nil {
			println(pNode.Val)
			stack.PushBack(pNode)
			pNode = pNode.Left
		} else {
			node := stack.Back().Value.(*TreeNode)
			pNode = node.Right
			stack.Remove(stack.Back())
		}
	}
}
