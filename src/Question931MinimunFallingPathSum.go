package main

import "math"

/**
https://leetcode.com/problems/minimum-falling-path-sum/
*/

func minimum(a []int) int {
	ans := math.MaxInt32
	for i := 0; i < len(a); i++ {
		if ans > a[i] {
			ans = a[i]
		}
	}
	return ans
}

func minFallingPathSum(A [][]int) int {
	result := make([][]int, len(A)+1)
	min := math.MaxInt32
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(A[0])+2)
		for j := 0; j < len(result[i]); j++ {
			if i == 0 {
				result[i][j] = 0
			} else if j == len(result[i])-1 || j == 0 {
				result[i][j] = math.MaxInt32
			} else {
				result[i][j] = minimum([]int{result[i-1][j-1], result[i-1][j], result[i-1][j+1]}) + A[i-1][j-1]
			}
		}
		if i == len(result)-1 {
			min = minimum(result[i])
		}
	}
	return min
}
