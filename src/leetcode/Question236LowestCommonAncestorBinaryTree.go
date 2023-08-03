package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func main() {
	root := constructTree([]string{"3", "5", "1", "6", "2", "0", "8", "null", "null", "7", "4"})
	p := TreeNode{5, nil, nil}
	q := TreeNode{4, nil, nil}
	lca := lowestCommonAncestor(root, &p, &q)
	print(lca.Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result, _, _ := lca(root, p, q)
	return result
}
func lca(root, p, q *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	var res_l, p_l, q_l, res_r, p_r, q_r *TreeNode
	if nil != root.Left {
		res_l, p_l, q_l = lca(root.Left, p, q)
		if nil != res_l {
			return res_l, p, q
		}
	}
	if nil != root.Right {
		res_r, p_r, q_r = lca(root.Right, p, q)
		if nil != res_r {
			return res_r, p, q
		}
	}
	var resP, resQ *TreeNode
	if p_l != nil || p_r != nil || root.Val == p.Val {
		resP = p
	}
	if q_l != nil || q_r != nil || root.Val == q.Val {
		resQ = q
	}
	if resP != nil && resQ != nil {
		return root, p, q
	} else {
		return nil, resP, resQ
	}
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
