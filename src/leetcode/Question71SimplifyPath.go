package main

import (
	"container/list"
	"strings"
)

/**
https://leetcode.com/problems/simplify-path/
*/
func simplifyPath(path string) string {
	values := strings.Split(path, "/")
	simples := list.New()
	for i := 0; i < len(values); i++ {
		if len(values[i]) == 0 || values[i] == "." {
			continue
		} else if values[i] == ".." && simples.Len() != 0 {
			simples.Remove(simples.Back())
		} else if values[i] != ".." {
			simples.PushBack(values[i])
		}
	}
	var result string
	for e := simples.Front(); e != nil; e = e.Next() {
		result += "/"
		result += e.Value.(string)
	}
	if len(result) == 0 {
		result = "/"
	}
	return result
}

func main() {
	//println(simplifyPath("/a//b////c/d//././/.."))
	println(simplifyPath("/../"))
}
