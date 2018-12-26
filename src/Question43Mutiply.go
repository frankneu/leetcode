package main

/**
https://leetcode.com/problems/multiply-strings/
*/
func multiply(num1 string, num2 string) string {
	result := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			result[len(num1)+len(num2)-i-j-2] += int((num1[i] - '0') * (num2[j] - '0'))
			result[len(num1)+len(num2)-i-j-1] += result[len(num1)+len(num2)-i-j-2] / 10
			result[len(num1)+len(num2)-i-j-2] %= 10
		}
	}
	var resultStr string
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != 0 {
			for j := i; j >= 0; j-- {
				resultStr += string(result[j] + '0')
			}
			return resultStr
		}
	}
	return "0"
}

func main() {
	println(multiply("12", "12"))
}
