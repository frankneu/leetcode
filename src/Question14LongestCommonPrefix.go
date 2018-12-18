package main

/**
https://leetcode.com/problems/longest-common-prefix/
*/

func longestCommonPrefix(strs []string) string {
	common := 0
	if 0 == len(strs) {
		return ""
	}
	for ; common < len(strs[0]); common++ {
		i := 1
		for ; i < len(strs); i++ {
			if common >= len(strs[i]) || strs[i][common] != strs[0][common] {
				break
			}
		}
		if i != len(strs) {
			break
		}
	}
	return strs[0][0:common]
}

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
