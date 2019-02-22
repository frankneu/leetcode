package main

import "container/list"

/**
https://leetcode.com/problems/leaf-similar-trees/
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generate(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	set := make(map[*TreeNode]int)
	result := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		node := queue.Back().Value.(*TreeNode)
		queue.Remove(queue.Back())
		if nil == node.Left && nil == node.Right {
			result = append(result, node.Val)
			continue
		}
		if nil != node.Left {
			if _, ok := set[node.Left]; ok {

			} else {
				set[node.Left] = 1
				queue.PushBack(node.Left)
			}
		}
		if nil != node.Right {
			if _, ok := set[node.Right]; ok {

			} else {
				set[node.Right] = 1
				queue.PushBack(node.Right)
			}
		}
	}
	return result
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	list1 := generate(root1)
	list2 := generate(root2)
	if len(list1) != len(list2) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func main() {
	node3 := new(TreeNode)
	node9 := new(TreeNode)
	node20 := new(TreeNode)
	node15 := new(TreeNode)
	node7 := new(TreeNode)
	node3.Val = 3
	node9.Val = 9
	node20.Val = 20
	node15.Val = 15
	node7.Val = 7
	node3.Left = node9
	node3.Right = node20
	node20.Left = node15
	node20.Right = node7
	value := generate(node3)
	for i := 0; i < len(value); i++ {
		println(value[i])
	}
}
