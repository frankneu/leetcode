package main

/**
https://leetcode.com/problems/permutations/
*/
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var result [][]int
	remain := nums[1:]
	value := permute(remain)
	var partValue [][]int
	for j := 0; j < len(value); j++ {
		for k := 0; k <= len(value[j]); k++ {
			front := make([]int, k)
			mid := make([]int, len(value[j])-k)
			back := make([]int, len(value[j])-k+1)
			source := make([]int, len(value[j])+1)
			copy(front, value[j][:k])
			copy(mid, value[j][k:])
			copy(back, append([]int{nums[0]}, mid...))
			copy(source, append(front, back...))
			partValue = append(partValue, source)
		}
	}
	result = append(result, partValue...)
	return result
}

func main() {
	nums := []int{2, 4, 6, 8}
	result := permute(nums)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			print(result[i][j], "	")
		}
		println("")
	}
}
