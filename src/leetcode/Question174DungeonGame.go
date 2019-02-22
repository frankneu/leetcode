package main

import "math"

/**
https://leetcode.com/problems/dungeon-game/
*/
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	m := len(dungeon)
	n := len(dungeon[0])
	maxHp := make([][]int, m+1)
	//对与边界不好控制的，可以考虑增加一列冗余的边界来拦截掉，避免大量的冗余判断
	for i := 0; i <= m; i++ {
		maxHp[i] = make([]int, n+1)
		if i == m {
			for j := 0; j <= n; j++ {
				maxHp[i][j] = math.MaxInt32
			}
		}
		maxHp[i][n] = math.MaxInt32
	}
	maxHp[m][n-1] = 1
	maxHp[m-1][n] = 1
	//由于最后一位的正负值会影响之前的路径选择，因此从后向前遍历搜索
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			need := min(maxHp[i+1][j], maxHp[i][j+1]) - dungeon[i][j]
			if need > 0 {
				maxHp[i][j] = need
			} else {
				maxHp[i][j] = 1
			}
		}
	}
	return maxHp[0][0]
}

func main() {
	println(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}), 7)
	println(calculateMinimumHP([][]int{{2}, {1}}), 1)
	println(calculateMinimumHP([][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}), 3)
}
