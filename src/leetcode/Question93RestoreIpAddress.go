package main

import (
	"container/list"
	"strconv"
	"strings"
)

/**
https://leetcode.com/problems/restore-ip-addresses/
*/
type IpType struct {
	Pre  []string
	Back string
}

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	if len(s) < 4 {
		return result
	}
	queue := list.New()
	root := &IpType{[]string{}, s}
	queue.PushBack(root)
	for queue.Len() != 0 {
		node := queue.Front().Value.(*IpType)
		queue.Remove(queue.Front())
		for i := 1; i <= 3 && i < len(node.Back); i++ {
			pre := make([]string, len(node.Pre))
			copy(pre, node.Pre)
			pre = append(pre, node.Back[:i])
			back := node.Back[i:]
			value1, _ := strconv.Atoi(pre[len(pre)-1])
			if value1 > 255 {
				continue
			}
			if value1 == 0 && len(pre[len(pre)-1]) != 1 {
				continue
			}
			if value1 != 0 && pre[len(pre)-1][0] == '0' {
				continue
			}
			if len(node.Pre) == 2 {
				value2, _ := strconv.Atoi(back)
				if value2 > 255 {
					continue
				} else if value2 == 0 && len(back) != 1 {
					continue
				} else if value2 != 0 && back[0] == '0' {
					continue
				} else {
					pre = append(pre, back)
					result = append(result, strings.Join(pre, "."))
				}
			} else {
				child := &IpType{pre, back}
				queue.PushBack(child)
			}
		}
	}
	return result
}

func main() {
	v := restoreIpAddresses("010010")
	//v := restoreIpAddresses("25525511135")
	for i := 0; i < len(v); i++ {
		println(v[i])
	}
}
