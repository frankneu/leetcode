package main

/**
https://leetcode.com/problems/search-in-rotated-sorted-array/
*/

func findPivotMax(nums []int) int {
	i := 0
	j := len(nums) - 1
	k := 0
	for i < j {
		k = (i + j) / 2
		if nums[k] > nums[j] && nums[k] > nums[i] {
			i = k
		} else if nums[k] < nums[j] && nums[k] < nums[i] {
			j = k
		} else {
			return k
		}
	}
	return k
}

/**
这个解法存在BUG
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	p := findPivotMax(nums)
	i := p + 1
	j := p
	if p == len(nums)-1 {
		i = 0
		j = len(nums) - 1
	}
	k := 0
	for i != j {
		if i > j {
			k = ((i+j)/2 + p) % len(nums)
		} else {
			k = (i + j) / 2
		}
		if nums[k] == target {
			return k
		} else if target < nums[k] {
			j = (k - 1) % len(nums)
		} else if target > nums[k] {
			i = (k + 1) % len(nums)
		}
	}
	return -1
}

/**
讨论中的一种巧妙算法，避免了求解具体转折点的麻烦
*/
func searchByOthers(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		k := (j-i)/2 + i
		if target == nums[k] {
			return k
		}
		//情况1. k在翻转点的左侧
		//情况2. k在翻转点右侧
		if nums[k] < nums[j] { //判断为情况1
			if nums[k] < target && nums[j] >= target {
				i = k + 1
			} else {
				j = k - 1
			}
		} else { //判断为情况2
			if nums[i] <= target && target < nums[k] {
				j = k - 1
			} else {
				i = k + 1
			}
		}
	}
	return -1
}

func main() {
	nums := []int{5, 1, 3}
	println(search(nums, 5))
}
