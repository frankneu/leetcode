package main

import "math"

/**
https://leetcode.com/problems/divide-two-integers/
I don't want to finish this boring question
*/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 0 {
		return 2147483647
	}
	i := 1
	if (divisor < 0 && dividend > 0) || (divisor > 0 && dividend < 0) {
		i = -1
	}
	e := float64(10)
	v := math.Log10(math.Abs(float64(dividend))) - math.Log10(math.Abs(float64(divisor)))
	result := int(math.Pow(e, v)) * i
	if result > 2147483647 || result < -2147483648 {
		return 2147483647
	} else {
		return result
	}
}

func main() {
	println(divide(2147483647, 1))
}
