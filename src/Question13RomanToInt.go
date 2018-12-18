package main

/**
https://leetcode.com/problems/roman-to-integer/
*/

var ROMAN_MAP = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func romanToInt(s string) int {
	result := 0
	i := 0
	for i < len(s) {
		if i < len(s)-1 && ((s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X')) || (s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C')) || (s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M'))) {
			result += ROMAN_MAP[string([]byte{s[i], s[i+1]})]
			i += 2
			continue
		}
		result += ROMAN_MAP[string([]byte{s[i]})]
		i++
	}
	return result
}

func main() {
	println(romanToInt("MCMXCIV"))
}
