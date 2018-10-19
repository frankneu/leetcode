package main

import (
	"fmt"
	"sort"
)

func main() {
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	//res = threeSum([]int{-2,0,0,2,2})
	for index, value := range res {
		fmt.Print(index, ":")
		for i, _ := range value {
			fmt.Print(res[index][i], " ")
		}
		fmt.Println()
	}
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	var index = 0
	for ; index < len(nums); index++ {
		var i, j = index + 1, len(nums) - 1
		if index > 0 && nums[index-1] == nums[index] {
			continue
		}
		if nums[index] > 0 {
			break
		}
		for i < j {
			b, c := nums[i], nums[j]
			if nums[j] < 0 {
				break
			}
			value := nums[index] + b + c
			if value == 0 {
				result = append(result, []int{nums[index], b, c})
			}

			if value >= 0 {
				for i < j && c == nums[j] && nums[j] >= 0 {
					j--
				}
			}
			if value <= 0 && nums[j] >= 0 {
				for i < j && b == nums[i] {
					i++
				}
			}
		}
	}
	return result
}

func threeSumOthers(nums []int) [][]int {
	length := len(nums)
	solution := make([][]int, 0)
	if length < 3 {
		return solution
	}
	sort.Ints(nums)
	for k, _ := range nums {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		l := k + 1
		r := length - 1
		for l < r {
			sum := nums[k] + nums[l] + nums[r]
			if sum == 0 {
				array := []int{nums[k], nums[l], nums[r]}
				solution = append(solution, array)
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}

		}
	}
	return solution
}
