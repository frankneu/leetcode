package main

import "math"

/**
https://leetcode.com/problems/minimum-size-subarray-sum/
1. 计算sum[i]为从0到i的参数之和
2. 遍历i，在i->end这段sum数组中，查找最小的sum[i]+target元素，由于sum是递增的，所以此处可以二分查找
*/
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := math.MaxInt32
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for i := 1; i < len(nums); i++ {
		target := s + sum[i-1]
		bound := lowerBound(target, sum[i:]) + i
		if bound != len(sum) && result > (bound-i) {
			result = bound - i + 1
		}
	}
	if result == math.MaxInt32 {
		result = 0
	}
	return result
}

func lowerBound(target int, sum []int) int {
	left := 0
	right := len(sum) - 1
	middle := (left + right) / 2
	for left <= right {
		if sum[middle] < target {
			left = middle + 1
		} else if sum[middle] > target {
			right = middle - 1
		} else if sum[middle] == target {
			break
		}
		middle = (left + right) / 2
	}
	if sum[middle] < target {
		middle++
	}
	return middle
}

func main() {
	println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
}
