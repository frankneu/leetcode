package main

import "strconv"

/**
https://leetcode.com/problems/count-and-say/
*/

func generate(word string) string {
	var result string
	count := 1
	current := word[0]
	for i := 1; i < len(word); i++ {
		if word[i] == current {
			count++
		} else {
			result += strconv.Itoa(count) + string(current)
			count = 1
			current = word[i]
		}
	}
	result += strconv.Itoa(count) + string(current)
	return result
}

func countAndSay(n int) string {
	start := "1"
	for i := 1; i < n; i++ {
		start = generate(start)
	}
	return start
}

func main() {
	println(countAndSay(10))
}
