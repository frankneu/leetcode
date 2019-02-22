package main

/**
https://leetcode.com/problems/unique-paths/
*/

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := m - 1; i >= 0; i-- {
		matrix[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			if i == m-1 || j == n-1 {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = matrix[i+1][j] + matrix[i][j+1]
			}
		}
	}
	return matrix[0][0]
}

func main() {
	println(uniquePaths(7, 3))
}
