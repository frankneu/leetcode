package main

import (
	"container/list"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/
*/
/**
获取所有父节点
*/
func findAncestors(root *TreeNode, target *TreeNode) []*TreeNode {
	if root.Val == target.Val {
		return []*TreeNode{root}
	}
	var left []*TreeNode
	var right []*TreeNode
	if nil != root.Left {
		left = findAncestors(root.Left, target)
	}
	if nil != root.Right {
		right = findAncestors(root.Right, target)
	}
	if len(left) != 0 {
		return append(left, root)
	} else if len(right) != 0 {
		return append(right, root)
	} else {
		return nil
	}
}

/**
进行深度为K的广度搜索，返回第K层数据
*/
func bfs(root *TreeNode, K int) []int {
	if nil == root {
		return nil
	}
	queue := list.New()
	next := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 && K != 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if node == nil {
			continue
		}
		if node.Left != nil {
			next.PushBack(node.Left)
		}
		if node.Right != nil {
			next.PushBack(node.Right)
		}
		if queue.Len() == 0 {
			K--
			queue = next
			next = list.New()
		}
	}
	result := make([]int, 0)
	for node := queue.Front(); node != nil; node = node.Next() {
		if node.Value != nil {
			result = append(result, node.Value.(*TreeNode).Val)
		}
	}
	return result
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	ancestor := findAncestors(root, target)
	var node *TreeNode
	var pre *TreeNode
	result := make([]int, 0)
	i := 0
	for ; i < K && i < len(ancestor); i++ {
		node = ancestor[i]
		var search *TreeNode
		length := 0
		//搜索另外一条边，当前搜索路径上不应该继续搜索
		if pre != nil {
			if node.Left != nil && node.Left.Val == pre.Val {
				search = node.Right
			} else if node.Right != nil && node.Right.Val == pre.Val {
				search = node.Left
			}
			length = K - i - 1
		} else {
			search = node
			length = K
		}
		result = append(result, bfs(search, length)...)
		pre = node
	}
	if i == K && i < len(ancestor) {
		result = append(result, ancestor[K].Val)
	}
	return result
}

/**
把leetcode树结构数组展开成树
*/
func constructTree(values []string) *TreeNode {
	nodeNumber := 1
	i := 0
	level := 1
	var curr []*TreeNode
	var pre []*TreeNode
	var root *TreeNode
	for i < len(values) {
		var currentValues []string
		if i+nodeNumber > len(values) {
			currentValues = values[i:]
		} else {
			currentValues = values[i : i+nodeNumber]
		}
		curr = make([]*TreeNode, nodeNumber)
		nodeNumber = 0
		index := 0
		for j := 0; j < len(currentValues); j++ {
			if "null" != currentValues[j] {
				node := new(TreeNode)
				curr[index] = node
				index++
				node.Val, _ = strconv.Atoi(currentValues[j])
				if j%2 == 0 && pre != nil {
					pre[j/2].Left = node
				} else if pre != nil {
					pre[j/2].Right = node
				}
				nodeNumber++
			}
		}
		nodeNumber *= 2
		pre = curr
		i += len(currentValues)
		if level == 1 {
			root = curr[0]
		}
		level++
	}
	return root
}

func main() {
	//tree := constructTree([]string{"3","5","1","6","2","0","8","null","null","7","4"})
	//tree := constructTree([]string{"0","1","3","null","2"})
	//tree := constructTree([]string{"0","2","1","null","null","3"})
	//tree := constructTree([]string{"0","5","1","null","null","2","6","null","3","null","null","4","null","7"})
	//tree := constructTree([]string{"0","null","1","null","2","null","3"})
	tree := constructTree([]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"})
	target := new(TreeNode)
	target.Val = 5
	result := distanceK(tree, target, 2)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}
