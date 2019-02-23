package main

import "math"

/**
https://www.hackerrank.com/challenges/2d-array/problem?filter=Northeastern+University%2C+Boston&filter_on=school&h_l=interview&page=1&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
*/

func hourglassSum(arr [][]int32) int32 {
	var ans int32 = math.MinInt32
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[0])-2; j++ {
			c := count(arr, i, j)
			if ans < c {
				ans = c
			}
		}
	}
	return ans
}

func count(arr [][]int32, i int, j int) int32 {
	var result int32
	result += arr[i][j] + arr[i][j+1] + arr[i][j+2]
	result += arr[i+1][j+1]
	result += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
	return result
}

func main() {
	arr := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	println(hourglassSum(arr), -6)
}
