package main

/**
https://leetcode.com/problems/split-array-largest-sum/
todo dp problem
*/
func splitArray(nums []int, m int) int {
	max := 0
	l := 0
	r := len(nums)
	for l <= r {
		m := (l + r) / 2
		left := 0
		for i := 0; i < m; i++ {
			left += nums[i]
		}
		right := 0
		for j := 0; j < len(nums); j++ {
			right += nums[j]
		}
		value := maxValue(left, right)

		if max < value {
			va
		}
	}
}

func maxValue(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func peakStatus() {

}
