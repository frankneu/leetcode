package main

/**
https://leetcode.com/problems/search-a-2d-matrix/
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	i := 0
	j := m*n - 1
	for i <= j {
		m = (i + j) / 2
		a := m / n
		b := m % n
		v := matrix[a][b]
		if v == target {
			return true
		} else if v < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	println(searchMatrix(matrix, 12))
}
