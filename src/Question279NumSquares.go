package main

/**
https://leetcode.com/problems/perfect-squares/
*/
func numSquares(n int) int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		for j := 1; j*j <= i; j++ {
			t := result[i-j*j] + 1
			if t < count {
				count = t
			} else if count == 0 {
				count = t
			}
		}
		result[i] = count
	}
	return result[n]
}

func main() {
	println(numSquares(12))
}
