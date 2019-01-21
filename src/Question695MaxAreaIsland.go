package main

/**
https://leetcode.com/problems/max-area-of-island/
*/
func countArea(grid [][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	result := 1
	grid[i][j] = 0
	result += countArea(grid, i+1, j)
	result += countArea(grid, i-1, j)
	result += countArea(grid, i, j+1)
	result += countArea(grid, i, j-1)
	return result
}

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				value := countArea(grid, i, j)
				if value > result {
					result = value
				}
				i = 0
				j = 0
			}
		}
	}
	return result
}

func main() {
	println(maxAreaOfIsland([][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}), 4)
}
