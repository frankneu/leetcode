package main

import (
	"math"
	"strconv"
	"strings"
)

func generateParenthesis(n int) []string {
	var result []string
	max := math.Pow(2, float64(n*2))
	min := math.Pow(2, float64(n*2-1))
	for min < max {
		value := strconv.FormatInt(int64(min), 2)
		if checkParenthesisMatch(value, n) {
			value = strings.Replace(value, "1", "(", -1)
			value = strings.Replace(value, "0", ")", -1)
			result = append(result, value)
		}
		min++
	}
	return result
}

func checkParenthesisMatch(value string, n int) bool {
	count := 0
	for i := 0; i < len(value); i++ {
		if count > n {
			return false
		}
		if count < 0 {
			return false
		}
		if 48 == int(value[i]) {
			count--
		}
		if 49 == int(value[i]) {
			count++
		}
	}
	return count == 0
}

func main() {
	//println(checkParenthesisMatch("111000",6))
	result := generateParenthesis(1)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}
