package main

func main() {
	print(convert("AB", 2))
}

/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "
PAHNAPLSIIGYIR
PAHNAPLSIIGYIR
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "
PINALSIGYAHRPI
PINALSIGYAHRPI
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
func convert(s string, numRows int) string {
	var store_str = make([]string, numRows)
	index := 0
	i := 0
	direct := true
	for ; i < len(s); i++ {
		store_str[index] += string(s[i])
		if numRows == 1 {
			continue
		}
		if index == 0 {
			index = 1
			direct = true
			continue
		}
		if index == numRows-1 {
			index = numRows - 2
			direct = false
			continue
		}
		if direct {
			index++
		} else {
			index--
		}
	}
	var result string
	for i = 0; i < numRows; i++ {
		result += store_str[i]
	}
	return result
}
