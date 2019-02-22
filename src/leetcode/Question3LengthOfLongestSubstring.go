package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(lengthOfLongestSubstringBetter2("abcabcbb"))
}

/**
Given a string, find the length of the longest substring without repeating characters.
Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

other test case
" "
"abba"
"tmmzuxt"
"abcabcbb"
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var maxLength int = 1
	var startIndex int = 0
	var curLength int = 1
	for i := 1; i < len(s); i++ {
		sub := string([]byte(s)[startIndex:i])
		target := string(s[i])
		index := strings.LastIndexAny(sub, target)
		if index >= 0 {
			startIndex += index + 1
			curLength = i - startIndex + 1
		} else {
			curLength++
		}
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}

func lengthOfLongestSubstringBetter(s string) int {
	//there is a better way, the map can be replaced by a array
	posMap := make(map[string]int)
	var maxLength = 0
	var curLength = 0
	var startIndex = 0
	for i := 0; i < len(s); i++ {
		if v, ok := posMap[string(s[i])]; ok {
			if v < startIndex {
				v = startIndex
			}
			curLength = i - v
			startIndex = v
		} else {
			curLength++
		}
		posMap[string(s[i])] = i
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}

func lengthOfLongestSubstringBetter2(s string) int {
	//there is a better way, the map can be replaced by a array
	var status [255]int
	for i := range status {
		status[i] = -1
	}
	var maxLength = 0
	var curLength = 0
	var startIndex = 0
	for i := 0; i < len(s); i++ {
		var v = status[s[i]]
		if v != -1 {
			if v < startIndex {
				v = startIndex
			}
			curLength = i - v
			startIndex = v
		} else {
			curLength++
		}
		status[s[i]] = i
		if curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}
