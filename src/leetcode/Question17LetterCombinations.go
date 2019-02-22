package main

import "strings"

var NUMBERS = map[int]string{
	50: "abc",
	51: "def",
	52: "ghi",
	53: "jkl",
	54: "mno",
	55: "pqrs",
	56: "tuv",
	57: "wxyz",
}

func letterCombinations(digits string) []string {
	var result []string
	var values []string
	var index = make([]int, len(digits))
	for i := 0; i < len(digits); i++ {

		values = append(values, NUMBERS[int(digits[i])])
	}
	var more = true
	for more {
		var bytes []string
		for i := 0; i < len(values); i++ {
			bytes = append(bytes, string(values[i][index[i]]))
		}
		word := strings.Join(bytes, "")
		if len(word) > 0 {
			result = append(result, word)
		}
		more = addOneIndex(index, values)
	}
	return result
}

func addOneIndex(index []int, values []string) bool {
	for i := len(index) - 1; i >= 0; i-- {
		if index[i]+1 == len(values[i]) {
			index[i] = 0
		} else {
			index[i] += 1
			return true
		}
	}
	return false
}

func main() {

	result := letterCombinations("")
	for i := 0; i < len(result); i++ {
		print(result[i], ",")
	}
}
