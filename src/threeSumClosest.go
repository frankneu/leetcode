package main

import (
	"math"
	"sort"
)

func main() {
	res := []int{-1, 2, 1, -4}
	println(threeSumClosest(res, 1))
}

func threeSumClosest(nums []int, target int) int {
	var index, i, j, result = 0, 1, len(nums) - 1, -target
	sort.Ints(nums)
	for index < len(nums) {
		i = index + 1
		j = len(nums) - 1
		for i < j {
			a, b, c := nums[index], nums[i], nums[j]
			value := a + b + c
			if value == target {
				return target
			}
			if math.Abs(float64(value-target)) < math.Abs(float64(result-target)) || target == 0 {
				result = value
			}
			if value > target {
				j--
			}
			if value < target {
				i++
			}
		}
		index++
	}
	return result
}
