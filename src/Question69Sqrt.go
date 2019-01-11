package main

/**
https://leetcode.com/problems/sqrtx/
*/

func mySqrt(x int) int {
	i := 0
	j := x
	m := 0
	for i <= j {
		m = (i + j) / 2
		squr := m * m
		if squr < x {
			i = m + 1
		} else if squr > x {
			j = m - 1
		} else {
			return m
		}
	}
	if m*m > x {
		m--
	}
	return m
}

func main() {
	println(mySqrt(8))
}
