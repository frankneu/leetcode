package main

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	m := 0
	for i <= j {
		m = (i + j) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	if nums[m] < target {
		m += 1
	}
	return m
}

func main() {
	nums := []int{1, 3, 5, 6}
	println(searchInsert(nums, 0))
}
