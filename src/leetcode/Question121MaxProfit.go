package main

/**
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0
	value := make([]int, len(prices))
	value[0] = 0
	for i := 1; i < len(prices); i++ {
		value[i] = max(prices[i]-prices[i-1]+value[i-1], 0)
		if result < value[i] {
			result = value[i]
		}
	}
	return result
}

func main() {
	println(maxProfit([]int{7, 6, 4, 3, 1}))
}
