package main

import "sort"

/**
https://leetcode.com/problems/combination-sum-ii/submissions/
*/
func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func singleCombinationSumRecursion(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	remain := target
	var result [][]int
	value1 := singleCombinationSumRecursion(candidates[1:], target)
	remain -= candidates[0]
	var value2 [][]int
	if remain == 0 {
		value2 = [][]int{{candidates[0]}}
	} else if remain > 0 {
		value2 = singleCombinationSumRecursion(candidates[1:], remain)
		for i := 0; i < len(value2); i++ {
			value2[i] = append(value2[i], candidates[0])
		}
	} else {
		return [][]int{}
	}
	result = append(result, value1...)
	return append(result, value2...)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	value := singleCombinationSumRecursion(candidates, target)
	for i := 0; i < len(value); i++ {
		j := 0
		for j < len(result) {
			if intSliceEqual(result[j], value[i]) {
				break
			} else {
				j++
			}
		}
		if j == len(result) {
			result = append(result, value[i])
		}
	}
	return result
}

func main() {
	nums := []int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34, 16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16, 8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12}
	result := combinationSum2(nums, 27)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			print(result[i][j], "	")
		}
		println("")
	}
}
