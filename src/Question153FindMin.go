package main

/**
注意终止条件，nums[i]<nums[j]这种情况，之后再进行中值搜索
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
*/
func findMin(nums []int) int {
	i := 0
	j := len(nums) - 1
	if j == -1 {
		return -1
	}
	m := 0
	for i < j {
		if nums[i] < nums[j] {
			return nums[i]
		}
		m = (i + j) / 2
		if nums[m] > nums[j] {
			i = m + 1
		} else {
			j = m
		}
	}
	return nums[i]
}

func main() {
	println(findMin([]int{4, 5, 6, 7, 0}))
}
