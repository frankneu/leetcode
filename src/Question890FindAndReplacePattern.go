package main

import (
	"strconv"
	"strings"
)

/**
https://leetcode.com/problems/find-and-replace-pattern/
*/

func generateIdf(pat string) map[string]byte {
	idf := make(map[byte][]string)
	for i := 0; i < len(pat); i++ {
		if v, ok := idf[pat[i]]; ok {
			v = append(v, strconv.Itoa(i))
			idf[pat[i]] = v
		} else {
			arr := make([]string, 0)
			arr = append(arr, strconv.Itoa(i))
			idf[pat[i]] = arr
		}
	}
	ans := make(map[string]byte)
	for key, value := range idf {
		ans[strings.Join(value, ",")] = key
	}
	return ans
}

func match(source, target map[string]byte) bool {
	if len(source) != len(target) {
		return false
	}
	for k, _ := range source {
		if _, ok := target[k]; ok {
			continue
		} else {
			return false
		}
	}
	return true
}

func findAndReplacePattern(words []string, pattern string) []string {
	patternMap := generateIdf(pattern)
	ans := make([]string, 0)
	for i := 0; i < len(words); i++ {
		wordMap := generateIdf(words[i])
		if match(wordMap, patternMap) {
			ans = append(ans, words[i])
		}
	}
	return ans
}

func main() {
	value := findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb")
	for i := 0; i < len(value); i++ {
		println(value[i])
	}
}
