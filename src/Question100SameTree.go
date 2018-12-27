package main

import "container/list"

/**
https://leetcode.com/problems/same-tree/
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	} else if (nil == p && nil != q) || (nil != p && nil == q) {
		return false
	}
	queuePre := list.New()
	queueQus := list.New()
	queuePre.PushBack(p)
	queueQus.PushBack(q)
	for queuePre.Len() > 0 && queueQus.Len() > 0 {
		eleP := (queuePre.Front().Value).(*TreeNode)
		queuePre.Remove(queuePre.Front())
		eleQ := (queueQus.Front().Value).(*TreeNode)
		queueQus.Remove(queueQus.Front())
		if eleP.Val != eleQ.Val {
			return false
		}
		if (nil == eleP.Left && nil != eleQ.Left) || (nil != eleP.Left && nil == eleQ.Left) {
			return false
		}
		if (nil == eleP.Right && nil != eleQ.Right) || (nil != eleP.Right && nil == eleQ.Right) {
			return false
		}
		if nil != eleP.Left {
			queuePre.PushBack(eleP.Left)
			queueQus.PushBack(eleQ.Left)
		}
		if nil != eleP.Right {
			queuePre.PushBack(eleP.Right)
			queueQus.PushBack(eleQ.Right)
		}
	}
	return queuePre.Len() == 0 && queueQus.Len() == 0
}
