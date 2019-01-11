package main

/**
https://leetcode.com/problems/add-binary/
*/
func addBinary(a string, b string) string {
	var result string
	x := '0'
	i := len(a) - 1
	j := len(b) - 1
	left := []rune(a)
	right := []rune(b)
	for i >= 0 || j >= 0 || x == '1' {
		y := '0'
		z := '0'
		if i >= 0 {
			y = left[i]
			i--
		}
		if j >= 0 {
			z = right[j]
			j--
		}
		if x == '1' && y == '1' && z == '1' {
			result = "1" + result
			x = '1'
		} else if (x == '1' && y == '1' && z != '1') || (x == '1' && y != '1' && z == '1') || (x != '1' && y == '1' && z == '1') {
			result = "0" + result
			x = '1'
		} else if (x == '1' && y != '1' && z != '1') || (x != '1' && y == '1' && z != '1') || (x != '1' && y != '1' && z == '1') {
			result = "1" + result
			x = '0'
		} else {
			result = "0" + result
			x = '0'
		}
	}
	return result
}

func main() {
	println(addBinary("1010", "1011"))
}
