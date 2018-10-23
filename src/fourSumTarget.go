package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	var min, max = 0, len(nums) - 1
	for min < max {
		var minV = nums[min]
		for min < max {
			var maxV = nums[max]
			i, j := min+1, max-1
			for i < j {
				var a, b = nums[i], nums[j]
				value := a + b + minV + maxV
				if value == target {
					array := []int{a, b, minV, maxV}
					result = append(result, array)
					i++
					j--
					for b == nums[j] && i < j {
						j--
					}
					for a == nums[i] && i < j {
						i++
					}
				} else if value > target {
					j--
				} else {
					i++
				}
			}
			for maxV == nums[max] && min < max {
				max--
			}
		}
		max = len(nums) - 1
		for minV == nums[min] && min < max {
			min++
		}
	}
	return result
}

func main() {
	res := fourSum([]int{-2, 0, -1, 0, 2, 1}, 0)
	for index, value := range res {
		fmt.Print(index, ":")
		for i, _ := range value {
			fmt.Print(res[index][i], " ")
		}
		fmt.Println()
	}
}
