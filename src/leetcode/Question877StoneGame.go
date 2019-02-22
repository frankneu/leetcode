package main

import "math"

/**
https://leetcode.com/problems/stone-game/
https://leetcode.com/problems/stone-game/discuss/154610/C%2B%2BJavaPython-DP-or-Just-return-true
*/
func stoneGameV1(piles []int) bool {
	//dp[i][j]存放的是从i->J这段石头中，先手玩家可以比后手玩家多拿的分数
	dp := make([][]float64, len(piles))
	for i := 0; i < len(piles); i++ {
		dp[i] = make([]float64, len(piles))
		dp[i][i] = float64(piles[i])
	}
	for j := 1; j < len(piles); j++ {
		for i := 0; i < len(piles)-j; i++ {
			//如果要拿第i或是第i+j位，上一步的优势就变成劣势，所以是当前拿到的分数减去上一步的分数
			dp[i][i+j] = math.Max(float64(piles[i])-dp[i+1][i+j], float64(piles[i+j])-dp[i][i+j-1])
		}
	}
	return dp[0][len(piles)-1] > 0
}

func stoneGame(piles []int) bool {
	//相当于只存储了上面DP矩阵中最后一列
	dp := make([]float64, len(piles))
	for d := 1; d < len(piles); d++ {
		for i := 0; i < len(piles)-d; i++ {
			dp[i] = math.Max(float64(piles[i])-dp[i+1], float64(piles[i+d])-dp[i])
		}
	}
	return dp[0] > 0
}
