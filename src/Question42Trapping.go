package main

/**
https://leetcode.com/problems/trapping-rain-water/
*/

func count(height []int, start int, end int) int {
	result := 0
	h := min(height[start], height[end])
	for i := start + 1; i < end; i++ {
		volume := h - height[i]
		if volume > 0 {
			result += volume
		}
	}
	return result
}

//从左到右的搜索，下面的方法不能很好的处理
func trapV1(height []int) int {
	if len(height) < 3 {
		return 0
	}
	start := 0
	end := 0
	result := 0
	cur := 0
	for start < len(height) && end < len(height) {
		if start == end && height[start] == 0 {
			start++
			end++
		} else if height[start] > 0 && start == end {
			end++
		} else if height[start] > 0 && height[end] >= height[start] {
			result += count(height, start, end)
			start = end
			cur = 0
		} else if height[start] > 0 && height[end] <= height[start] {
			tmp := count(height, start, end)
			if tmp >= cur {
				cur = tmp
			} else {
				result += cur
				cur = 0
				start = end - 1
			}
			end++
		} else {
			println(start, end)
		}
	}
	return result + cur
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

/**
对每个水位，检查其能够获取的最高水位，算法先获取左边和右边的最大值，然后算出结果
*/
func trapV2(height []int) int {
	result := 0
	for i := 1; i < len(height)-1; i++ {
		left := 0
		right := 0
		for j := 0; j <= i; j++ {
			left = max(left, height[j])
		}
		for j := i; j < len(height); j++ {
			right = max(right, height[j])
		}
		result += min(left, right) - height[i]
	}
	return result
}

/**
动态规划，存储左右的最大值
*/
func trap(height []int) int {
	if 0 == len(height) {
		return 0
	}
	leftMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax := make([]int, len(height))
	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	result := 0
	for i := 1; i < len(height)-1; i++ {
		result += min(leftMax[i], rightMax[i]) - height[i]
	}
	return result
}

func main() {
	println(trap([]int{9, 6, 8, 8, 5, 6, 3}), 3)
	println(trap([]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3}), 23)
	println(trap([]int{5, 4, 1, 2}), 1)
	println(trap([]int{2, 0, 2}), 2)
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
}
