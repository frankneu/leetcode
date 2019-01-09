package main

/**
https://leetcode.com/problems/length-of-last-word/
*/
func lengthOfLastWord(s string) int {
	var result int
	for i := len(s) - 1; i > -1; i-- {
		if ' ' != s[i] {
			result++
		} else if (i < len(s)-1) && s[i+1] != ' ' {
			break
		}
	}
	return result
}
