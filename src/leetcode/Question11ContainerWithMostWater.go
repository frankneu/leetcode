package main

func main() {
	//p:=[]int {30,35,15,5,10,20,15}
	p := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	println(maxAreaBetter(p))
}

/**
https://leetcode.com/problems/container-with-most-water/
*/
func maxArea(height []int) int {
	var i, j, result int = 1, 0, 0
	for ; i < len(height); i++ {
		j = 0
		for ; j < i; j++ {
			value := (i - j) * min(height[i], height[j])
			if value > result {
				result = value
			}

		}
	}
	return result
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxAreaBetter(height []int) int {
	var i, j, result = 0, len(height) - 1, 0
	for i < j {
		value := (j - i) * min(height[i], height[j])
		if value > result {
			result = value
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}

	}
	return result
}
