package main

/**
https://leetcode.com/problems/text-justification/
*/
func fullJustify(words []string, maxWidth int) []string {
	result := make([][]string, 0)
	current := make([]string, 0)
	flag := false
	for i := 0; i < len(words); {
		now := make([]string, len(current))
		copy(now, current)
		now = append(now, words[i])
		length := length(now, true)
		if length > maxWidth {
			result = append(result, current)
			current = make([]string, 0)
			flag = true
		} else if length == maxWidth {
			result = append(result, now)
			current = make([]string, 0)
			flag = true
			i++
		} else {
			current = now
			flag = false
			i++
		}
	}
	if !flag {
		result = append(result, current)
	}
	ans := make([]string, 0)
	for i := 0; i < len(result); i++ {
		if len(result[i]) == 1 {
			ans = append(ans, appendSpace(result[i][0], maxWidth-len(result[i][0])))
		} else if i != len(result)-1 {
			space := maxWidth - length(result[i], false)
			width := space / (len(result[i]) - 1)
			count := len(result[i]) - 1
			if space%(len(result[i])-1) != 0 {
				width++
				count = space - (len(result[i])-1)*(width-1)
			}
			var value string
			for j := 0; j < len(result[i])-1; j++ {
				if count > 0 {
					value += appendSpace(result[i][j], width)
					count--
				} else {
					value += appendSpace(result[i][j], space/(len(result[i])-1))
				}
			}
			value += result[i][len(result[i])-1]
			ans = append(ans, value)
		} else {
			var value string
			for j := 0; j < len(result[i])-1; j++ {
				value += appendSpace(result[i][j], 1)
			}
			value += result[i][len(result[i])-1]
			value = appendSpace(value, maxWidth-len(value))
			ans = append(ans, value)
		}
	}
	return ans
}

func length(words []string, with bool) int {
	length := 0
	for i := 0; i < len(words); i++ {
		length += len(words[i])
	}
	if with {
		length += len(words) - 1
	}
	return length
}

func appendSpace(word string, n int) string {
	var space string
	for i := 0; i < n; i++ {
		space += " "
	}
	return word + space
}

func main() {
	ans := fullJustify([]string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}, 16)
	for i := 0; i < len(ans); i++ {
		println(len(ans[i]))
	}
}
