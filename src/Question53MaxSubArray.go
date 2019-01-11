package main

/**
https://leetcode.com/problems/maximum-subarray/
*/

func maxSubArray(nums []int) int {
	value := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}
	value[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if value[i-1] > 0 {
			value[i] = value[i-1] + nums[i]
		} else {
			value[i] = nums[i]
		}
		if ans < value[i] {
			ans = value[i]
		}
	}
	return ans
}
