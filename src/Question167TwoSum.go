package main

/**
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
*/

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		value := target - numbers[i]
		k := len(numbers) - 1
		j := i + 1
		for j <= k {
			m := (j + k) / 2
			if numbers[m] == value {
				return []int{i + 1, m + 1}
			} else if numbers[m] > value {
				k = m - 1
			} else {
				j = m + 1
			}
		}
	}
	return []int{}
}

func main() {
	value := twoSum([]int{2, 7, 11, 15}, 9)
	println(value[0])
	println(value[1])
}
