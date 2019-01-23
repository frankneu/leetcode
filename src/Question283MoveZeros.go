package main

/**
https://leetcode.com/problems/move-zeroes/
*/
/**
使用前后两个指针
*/
func moveZeroesV1(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < right {
		for left < len(nums) && nums[left] != 0 {
			left++
		}
		for right >= 0 && nums[right] == 0 {
			right--
		}
		if left < len(nums) && right >= 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
}

/**
只使用一个指针
*/
func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i-count] = nums[i]
		} else {
			count++
		}
	}
	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 1, 2, 0, 0, 1, 2, 3, 0, 0}
	moveZeroes(nums)
	println(nums)
}
