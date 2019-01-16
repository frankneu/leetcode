package main

import (
	"strconv"
	"strings"
)

/**
https://leetcode.com/problems/complex-number-multiplication/
*/
func complexNumberMultiply(a string, b string) string {
	left := strings.Split(a, "+")
	right := strings.Split(b, "+")
	realLeft := 0
	imaginaryLeft := 0
	realRight := 0
	imaginaryRight := 0
	if len(left) == 2 {
		v, _ := strconv.Atoi(left[0])
		realLeft = v
		r, _ := strconv.Atoi(left[1][:len(left[1])-1])
		imaginaryLeft = r
	} else if len(left) == 1 {
		if left[0][len(left[0])-1] == 'i' {
			r, _ := strconv.Atoi(left[0][:len(left[0])-1])
			imaginaryLeft = r
		} else {
			v, _ := strconv.Atoi(left[0])
			realLeft = v
		}
	}

	if len(right) == 2 {
		v, _ := strconv.Atoi(right[0])
		realRight = v
		r, _ := strconv.Atoi(right[1][:len(right[1])-1])
		imaginaryRight = r
	} else if len(right) == 1 {
		if right[0][len(right[0])-1] == 'i' {
			r, _ := strconv.Atoi(right[0][:len(right[0])-1])
			imaginaryRight = r
		} else {
			v, _ := strconv.Atoi(right[0])
			realRight = v
		}
	}
	realAns := realLeft*realRight - imaginaryLeft*imaginaryRight
	imaginaryAns := realLeft*imaginaryRight + imaginaryLeft*realRight
	return strconv.Itoa(realAns) + "+" + strconv.Itoa(imaginaryAns) + "i"
}

func main() {
	println(complexNumberMultiply("1+-1i", "1+-1i"))
}
