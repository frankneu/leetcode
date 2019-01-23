package main

/**
https://leetcode.com/problems/valid-anagram/
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	MS := make(map[uint8]int)
	TS := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		MS[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		TS[t[i]]++
	}
	for k, v := range MS {
		if v != TS[k] {
			return false
		}
	}
	return true
}

func main() {
	println(isAnagram("anagram", "nagaram"), true)
	println(isAnagram("rat", "car"), false)
}
