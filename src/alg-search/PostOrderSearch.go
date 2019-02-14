package main

import "container/list"

/**
第一种思路：
对于任一结点P，将其入栈，然后沿其左子树一直往下搜索，直到搜索到没有左孩子的结点，此时该结点出现在栈顶，但是此时不能将其出栈并访问，因此其右孩子还为被访问。
所以接下来按照相同的规则对其右子树进行相同的处理，当访问完其右孩子时，该结点又出现在栈顶，此时可以将其出栈并访问。
这样就保证了正确的访问顺序。
可以看出，在这个过程中，每个结点都两次出现在栈顶，只有在第二次出现在栈顶时，才能访问它。
因此需要多设置一个变量标识该结点是否是第一次出现在栈顶。

第二种思路：
要保证根结点在左孩子和右孩子访问之后才能访问，因此对于任一结点P，先将其入栈。
如果P不存在左孩子和右孩子，则可以直接访问它；或者P存在左孩子或者右孩子，但是其左孩子和右孩子都已被访问过了，则同样可以直接访问该结点。
若非上述两种情况，则将P的右孩子和左孩子依次入栈，这样就保证了每次取栈顶元素的时候，左孩子在右孩子前面被访问，左孩子和右孩子都在根结点前面被访问。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrderTraverse2(root *TreeNode) {
	stack := list.New()
	var cur *TreeNode
	var pre *TreeNode
	stack.PushBack(root)
	for stack.Len() != 0 {
		cur = stack.Back().Value.(*TreeNode)
		if (cur.Left == nil && cur.Right == nil) || (pre != nil && (pre == cur.Left || pre == cur.Right)) {
			//上一次访问的节点不为空，而且这个节点是当前节点的子节点，这时候访问到当前节点只有一个可能，就是子节点都被访问过了
			println(cur.Val)
			stack.Remove(stack.Back())
			pre = cur
		} else {
			if cur.Right != nil {
				stack.PushBack(cur.Right)
			}
			if cur.Left != nil {
				stack.PushBack(cur.Left)
			}
		}
	}
}
