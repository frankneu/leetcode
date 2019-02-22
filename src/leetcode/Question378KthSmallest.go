package main

/**
https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/
*/
//二分查找
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	lo := matrix[0][0]
	hi := matrix[n-1][n-1]
	var mid, count int
	for lo < hi {
		mid = lo + (hi-lo)/2
		count = countLEQ(matrix, mid)
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

//计算矩阵中在X左侧的元素数
func countLEQ(matrix [][]int, x int) int {
	n := len(matrix)
	count := 0
	var j int
	for _, row := range matrix {
		for j = 0; j < n && row[j] <= x; j++ {
		}
		count += j
	}
	return count
}
