package main

/**
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
*/
func maxProfit(prices []int) int {
	ans := 0
	if len(prices) < 2 {
		return ans
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}), 7)
	println(maxProfit([]int{1, 2, 3, 4, 5}), 4)
	println(maxProfit([]int{7, 6, 4, 3, 1}), 0)
}
