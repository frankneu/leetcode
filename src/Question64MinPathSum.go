package main

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

/**
https://leetcode.com/problems/minimum-path-sum/
*/
func minPathSum(grid [][]int) int {
	result := make([][]int, len(grid))
	for i := len(grid) - 1; i >= 0; i-- {
		result[i] = make([]int, len(grid[i]))
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j == len(grid[i])-1 {
				result[i][j] = grid[i][j]
				continue
			}
			if i == len(grid)-1 {
				result[i][j] = grid[i][j] + result[i][j+1]
			} else if j == len(grid[i])-1 {
				result[i][j] = grid[i][j] + result[i+1][j]
			} else {
				result[i][j] = grid[i][j] + min(result[i][j+1], result[i+1][j])
			}
		}
	}
	return result[0][0]
}
