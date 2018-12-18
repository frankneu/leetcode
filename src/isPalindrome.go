package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	result := 0
	for y > 0 {
		r := y % 10
		result *= 10
		result += r
		y /= 10
	}
	return result == x
}

func main() {
	println(isPalindrome(100))
}
