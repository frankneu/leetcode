package main

func main() {}

/**
https://leetcode.com/problems/regular-expression-matching/
*/
func isMatch(s string, p string) bool {
	var i, j = 0, 0
	var match string
	var repeat = false
	var a, b uint8
	for ; i < len(s) && j < len(p); i++ {
		a, b = s[i], p[j]
		if '.' == b {
			b = a
			match += string(b)
			j++
			continue
		}
		if '*' == b {
			if j == 0 {
				return false
			}
			repeat = true
			b = match[j-1]
			continue
		} else {
			repeat = false
		}
		if a == b {
			match += string(b)

		}
	}
	return true
}
