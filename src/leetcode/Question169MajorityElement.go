package main

/**
https://leetcode.com/problems/majority-element/
*/
func majorityElement(nums []int) int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		if count[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}

func main() {
	println(majorityElement([]int{3, 2, 3}), 3)
	println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}), 2)
}
