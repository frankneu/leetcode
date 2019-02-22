package main

import "strconv"

/**
https://leetcode.com/problems/binary-tree-paths/
*/

func binaryTreePaths(root *TreeNode) []string {
	if nil == root {
		return nil
	}
	if nil == root.Left && nil == root.Right {
		return []string{strconv.Itoa(root.Val)}
	}
	result := make([]string, 0)
	if nil != root.Left {
		result = append(result, binaryTreePaths(root.Left)...)
	}
	if nil != root.Right {
		result = append(result, binaryTreePaths(root.Right)...)
	}
	for i := 0; i < len(result); i++ {
		result[i] = strconv.Itoa(root.Val) + "->" + result[i]
	}
	return result
}
