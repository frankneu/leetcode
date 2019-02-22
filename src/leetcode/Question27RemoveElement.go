package main

/**
https://leetcode.com/problems/remove-element/
*/

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for ; j < len(nums); j++ {
		if val != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	result := removeElement(nums, 2)
	for i := 0; i < result; i++ {
		println(nums[i])
	}
}
