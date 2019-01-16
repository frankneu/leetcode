package main

/**
https://leetcode.com/problems/counting-bits/
*/

func countBits(num int) []int {
	result := make([]int, num+1)
	result[0] = 0
	if num == 0 {
		return result
	}
	result[1] = 1
	if num == 1 {
		return result
	}
	start := 2
	for i := start; i <= num; i++ {
		if i == start {
			result[i] = 1
		} else if i == start*2 {
			start *= 2
			result[i] = 1
		} else {
			result[i] = 1 + result[i-start]
		}
	}
	return result
}

func main() {
	v := countBits(5)
	for i := 0; i < len(v); i++ {
		println(v[i])
	}
}
