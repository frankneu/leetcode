package main

/**
https://leetcode.com/problems/unique-binary-search-trees/
二分查找树是有序的
*/
func numTrees(n int) int {
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			result[i] += result[j-1] * result[i-j]
		}
	}
	return result[n]
}

func main() {
	println(numTrees(3))
}
