package main

import (
	"container/list"
)

/**
https://leetcode.com/problems/valid-parentheses/
*/

func isValid(s string) bool {
	stack := list.New()
	i := 0
	for ; i < len(s); i++ {
		if 0 == stack.Len() {
			stack.PushBack(s[i])
		} else {
			var back = stack.Back()
			//通过反射，指定interface对应的原始类型
			var v = back.Value.(uint8)
			var u = s[i]
			if (v == 40 && u == 41) || (v == 91 && u == 93) || (v == 123 && u == 125) {
				stack.Remove(stack.Back())
			} else {
				stack.PushBack(s[i])
			}
		}
	}
	return 0 == stack.Len()
}

func main() {
	println(isValid("([)]"))

}
