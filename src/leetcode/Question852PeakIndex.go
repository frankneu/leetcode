package main

/**
https://leetcode.com/problems/peak-index-in-a-mountain-array/
*/
func peakIndexInMountainArray(A []int) int {
	i := 0
	j := len(A) - 1
	m := (i + j) / 2
	for i <= j {
		if A[m] > A[i] && A[m] < A[j] {
			i = m + 1
		} else if A[m] < A[i] && A[m] > A[j] {
			j = m - 1
		} else if A[m] > A[m-1] && A[m] > A[m+1] {
			return m
		} else {
			i++
		}
		m = (i + j) / 2
	}
	return m
}

func main() {
	println(peakIndexInMountainArray([]int{3, 4, 5, 0}))
}
