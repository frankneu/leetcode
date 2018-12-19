package main

/**
https://leetcode.com/problems/remove-duplicates-from-sorted-array
*/

func removeDuplicates(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	value := nums[0]
	i := 1
	j := 1
	for ; j < len(nums); j++ {
		if nums[j] != value {
			nums[i] = nums[j]
			value = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{1, 1, 2}
	result := removeDuplicates(nums)
	for i := 0; i < result; i++ {
		println(nums[i])
	}
}
