package main

func main() {
	print(longestPalindrome("cbbd"))
}

/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
func longestPalindrome(s string) string {
	var store_str = make([]string, len(s))
	var result string
	for i := 0; i < len(s); i++ {
		var palindromic string
		if i-1 >= 0 {
			palindromic = store_str[i-1]
		}
		var ch1, ch2 uint8
		if i-len(palindromic)-1 >= 0 && len(palindromic) > 0 {
			ch1 = s[i-len(palindromic)-1]
		}
		ch2 = s[i]
		var newp string
		if ch1 == ch2 {
			store_str[i] = string([]byte(s)[i-len(palindromic)-1 : i+1])
			newp = store_str[i]
		} else {
			newp = findLongestPalindrome(palindromic, string(ch2))
			store_str[i] = newp
		}
		if len(newp) > len(result) {
			result = newp
		}
	}
	return result
}

func findLongestPalindrome(s string, ch string) string {
	i := 0
	for ; i < len(s); i++ {
		source := string([]byte(s)[i:len(s)]) + ch
		if checkPalindrome(source) {
			return source
		}
	}
	return ch
}

func checkPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
