package main

/**
https://leetcode.com/problems/implement-strstr/
use KMP algorithm
*/

func next(needle string) []int {
	next := make([]int, len(needle))
	j := 0
	i := 1
	for i < len(needle) {
		if needle[i] == needle[j] {
			next[i] = j + 1
			j++
			i++
		} else {
			if j != 0 {
				j = next[j-1]
			} else {
				next[i] = 0
				i++
			}
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := next(needle)
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j != 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	if j == len(needle) {
		return i - len(needle) + 1
	}
	return -1
}

func main() {
	println(strStr("aaaaa", "bba"))

}
