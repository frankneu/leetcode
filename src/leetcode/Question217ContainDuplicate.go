package main

/**
https://leetcode.com/problems/contains-duplicate/
*/
func containsDuplicate(nums []int) bool {
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		set[nums[i]]++
		if set[nums[i]] > 1 {
			return true
		}
	}
	return false
}

func main() {
	println(containsDuplicate([]int{1, 2, 3, 4, 1}), true)
	println(containsDuplicate([]int{1, 2, 3, 4, 5}), false)
}
