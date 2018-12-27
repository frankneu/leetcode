package main

/**
https://leetcode.com/problems/house-robber/
*/
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	result := 0
	resultArray := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			resultArray[i] = nums[i]
		} else if i == 1 {
			resultArray[i] = max(resultArray[i-1], nums[i])
		} else {
			resultArray[i] = max(resultArray[i-1], nums[i]+resultArray[i-2])
		}
		if resultArray[i] > result {
			result = resultArray[i]
		}
	}
	return result
}

func main() {
	nums := []int{1, 4, 5, 3, 2, 2}
	println(rob(nums))
}
