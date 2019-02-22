package main

import "strconv"

/**
https://leetcode.com/problems/decode-ways/
*/
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		var sub1 string
		var sub2 string
		if i == 0 {
			sub1 = s[i : i+1]
		} else {
			sub1 = s[i : i+1]
			sub2 = s[i-1 : i+1]
		}
		int1, _ := strconv.ParseInt(sub1, 10, 8)
		int2, _ := strconv.ParseInt(sub2, 10, 8)
		i2 := 0
		if i == 0 && int1 != 0 {
			dp[i] = 1
			continue
		}
		if i < 2 {
			i2 = 1
		} else {
			i2 = dp[i-2]
		}
		if int2 < 27 && int2 > 10 && int1 != 0 {
			dp[i] = dp[i-1] + i2
		} else if (int2 < 10 && int2 > 0) || (int2 > 26 && int1 != 0) {
			dp[i] = dp[i-1]
		} else if int2 > 0 && int2 < 27 && int1 == 0 {
			dp[i] = i2
		} else if (len(sub2) > 0 && int2 == 0) && (int2 > 26 && int1 == 0) {
			return 0
		}
	}
	return dp[len(s)-1]
}

func main() {
	println(numDecodings("12120"), 3)
	println(numDecodings("27"), 1)
	println(numDecodings("0"), 0)
	println(numDecodings("110"), 1)
	println(numDecodings("230"), 0)
	println(numDecodings("12210"), 3)
	println(numDecodings("12"), 2)
	println(numDecodings("226"), 3)
	println(numDecodings("01"), 0)
	println(numDecodings("10"), 1)

}
