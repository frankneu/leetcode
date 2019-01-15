package main

/**
https://leetcode.com/problems/implement-strstr/
use KMP algorithm
*/
func next(needle string) []int {
	//next数组的含义就是这个位置紧邻的前面有几个字符和开头的字符是一致的
	next := make([]int, len(needle))
	j := 0
	i := 1
	for i < len(needle) {
		if needle[i] == needle[j] {
			//当前位置与开头的一致，next的值向上累加
			next[i] = j + 1
			j++
			i++
		} else {
			if j != 0 {
				//不匹配的位置不是开头，读取之前一个节点的偏移，尝试向后退，让这个不匹配的字节匹配之前的一个节点，这部分有可能重复进入，如果是一个跟开头部分完全不一致的节点，最终会进到下面的循环
				j = next[j-1]
			} else {
				//这个字节彻底匹配不了，i进入下一个字节
				next[i] = 0
				i++
			}
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := next(needle)
	i := 0 //i是在目标字符串上走过的位置
	j := 0 //j是在要去寻找的字符串上走的位置
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			//不匹配的时候，如果在寻找的字符串上不是在开头，那就可以快速的读next数组中的偏移量
			if j != 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	if j == len(needle) {
		return i - len(needle) + 1
	}
	return -1
}

func main() {
	println(strStr("aaaaa", "bba"))

}
