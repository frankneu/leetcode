package main

/**
https://leetcode.com/problems/unique-paths-ii/
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	result := make([][]int, len(obstacleGrid))
	for i := len(obstacleGrid) - 1; i >= 0; i-- {
		if len(obstacleGrid[i]) == 0 {
			return 0
		}
		result[i] = make([]int, len(obstacleGrid[i]))
		for j := len(obstacleGrid[i]) - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
				continue
			}
			if i == len(obstacleGrid)-1 && j == len(obstacleGrid[i])-1 {
				result[i][j] = 1
				continue
			}
			if (i == len(obstacleGrid)-1 && result[i][j+1] == 0) || (j == len(obstacleGrid[i])-1 && result[i+1][j] == 0) {
				result[i][j] = 0
			} else if i == len(obstacleGrid)-1 || j == len(obstacleGrid[i])-1 {
				result[i][j] = 1
			} else {
				result[i][j] = result[i][j+1] + result[i+1][j]
			}
		}
	}
	return result[0][0]
}

func main() {
	obs := [][]int{{0}, {0}}
	println(uniquePathsWithObstacles(obs))
}
