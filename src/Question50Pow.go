package main

/**
https://leetcode.com/problems/powx-n/
*/
func myPow(x float64, n int) float64 {
	if x == 1 {
		return 1
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n == 1 {
		return x
	}
	value := myPow(x, n/2)
	if n%2 == 0 {
		return value * value
	} else {
		return value * value * x
	}
}

func main() {
	println(myPow(2, 10))
}
