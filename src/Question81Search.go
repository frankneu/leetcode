package main

/**
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
*/
func search(nums []int, target int) bool {
	i := 0
	if len(nums) == 0 {
		return false
	}
	j := len(nums) - 1
	m := 0
	for i <= j {
		m = (i + j) / 2
		if nums[m] == target {
			return true
		}
		if nums[i] == nums[j] {
			i++
			continue
		}
		if (nums[m] < target && nums[m] > nums[j]) || (nums[m] > nums[j] && target <= nums[j]) || (target > nums[m] && target <= nums[j] && nums[m] < nums[j]) {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return nums[m] == target
}

func main() {
	println(search([]int{1, 1, 3, 1}, 3))
}
