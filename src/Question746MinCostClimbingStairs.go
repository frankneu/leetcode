package main

/**
https://leetcode.com/problems/min-cost-climbing-stairs/
*/
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	value := make([]int, len(cost))
	value[0] = cost[0]
	value[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		value[i] = min(value[i-1], value[i-2]) + cost[i]
	}
	return min(value[len(cost)-1], value[len(cost)-2])
}
