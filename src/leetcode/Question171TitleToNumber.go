package main

/**
https://leetcode.com/problems/excel-sheet-column-number/
*/
func titleToNumber(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		v := s[i] - 'A' + 1
		ans *= 26
		ans += int(v)
	}
	return ans
}

func main() {
	println(titleToNumber("A"), 1)
	println(titleToNumber("AB"), 28)
	println(titleToNumber("ZY"), 701)
}
