package main

/**
https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	i := 0
	for i < len(nums) {
		v, ok := m[target-nums[i]]
		if ok {
			return []int{v, i}
		} else {
			m[nums[i]] = i
		}
		i++
	}
	return []int{}
}

func main() {
	index := []int{2, 9, 7, 11, 15}
	target := 9
	result := twoSum(index, target)
	i := 0
	for ; i < len(result); i++ {
		println(result[i])
	}

}
