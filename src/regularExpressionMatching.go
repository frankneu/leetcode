package main

func main() {
	println(isMatch("ab", "a*"))
	println(isMatch("ab", ".*"))
	println(isMatch("ab", "a"))
	println(isMatch("aab", "c*a*b"))
	println(isMatch("mississippi", "mis*is*p*."))
}

/**
https://leetcode.com/problems/regular-expression-matching/

https://leetcode.com/problems/regular-expression-matching/solution/
*/
func isMatch(s string, p string) bool {
	var i, j = 0, 0
	var match string
	var repeat = false
	var a, b uint8 = s[i], p[j]
	for ; i < len(s); i++ {
		a = s[i]
		if '.' == b {
			b = a
			match += string(b)
			j++
			if j < len(p) {
				b = p[j]
			}
			continue
		}
		if '*' == b {
			if j == 0 {
				return false
			}
			repeat = true
			b = match[j-1]
			i--
			continue
		} else {
			repeat = false
		}
		if a == b {
			match += string(b)
			if repeat {
				continue
			} else {
				j++
				if j < len(p) {
					b = p[j]
				}
			}
		} else {
			i--
			j++
			if j < len(p) {
				b = p[j]
			}
		}
	}
	if i < len(s) {
		return false
	} else {
		return true
	}
}
