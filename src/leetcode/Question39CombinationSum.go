package main

import "sort"

/**
https://leetcode.com/problems/combination-sum/
*/

func combinationSumRecursion(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	var result [][]int
	remain := target
	var set []int
	for remain >= candidates[0] {
		value := combinationSumRecursion(candidates[1:], remain)
		for i := 0; i < len(value); i++ {
			value[i] = append(value[i], set...)
		}
		result = append(result, value...)
		set = append(set, candidates[0])
		remain -= candidates[0]
		if remain == 0 {
			result = append(result, set)
			continue
		}
	}
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSumRecursion(candidates, target)
}
func main() {
	nums := []int{3, 4, 7, 8}
	//sort.Sort(IntSlice(a))
	result := combinationSum(nums, 11)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			print(result[i][j], "	")
		}
		println("")
	}
}
