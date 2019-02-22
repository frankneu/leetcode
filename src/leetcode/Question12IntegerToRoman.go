package main

func main() {
	println(intToRoman(3))
	println(intToRoman(4))
	println(intToRoman(9))
	println(intToRoman(58))
	println(intToRoman(1994))
	println(intToRoman(3985))
	println(intToRoman(40))
}

var NUMBER = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var ROMAN = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	var result string
	var i = 0
	for num > 0 {
		if num >= NUMBER[i] {
			result += ROMAN[i]
			num -= NUMBER[i]
		} else {
			i++
		}
	}
	return result
}
