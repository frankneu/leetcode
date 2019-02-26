package main

/**
https://leetcode.com/problems/max-increase-to-keep-city-skyline/
*/
func maxIncreaseKeepingSkyline(grid [][]int) int {
	tb := make([]int, len(grid))
	lr := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if tb[i] < grid[i][j] {
				tb[i] = grid[i][j]
			}
			if lr[j] < grid[i][j] {
				lr[j] = grid[i][j]
			}
		}
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if tb[i] > lr[j] {
				count += lr[j] - grid[i][j]
			} else {
				count += tb[i] - grid[i][j]
			}
		}
	}
	return count
}
