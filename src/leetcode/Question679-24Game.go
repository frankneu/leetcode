package main

import "errors"

/**
https://leetcode.com/problems/24-game/
递归 dfs
减法都大剪小，除法也是
todo
*/
func judgePoint24(nums []int) bool {
	_, e := compute(nums)
	return e != nil
}

func compute(nums []int) ([]int, error) {
	source, remain := selectTwo(nums)
	if len(nums) == 2 {
		res := operator(source[0])
		for i := 0; i < len(res); i++ {
			if res[i] == 24 {
				return nil, errors.New("Get")
			}
		}
		return res, nil
	}
	ans := make([]int, 0)
	for i := 0; i < len(source); i++ {
		s := source[i]
		r := remain[i]
		v := operator(s)
		for j := 0; j < len(v); j++ {
			n := make([]int, len(r))
			copy(n, r)
			n = append(n, v[j])
			sub, error := compute(n)
			if error != nil {
				return nil, error
			}
			ans = append(ans, sub...)
		}
	}
	return ans, nil
}

func selectTwo(nums []int) ([][]int, [][]int) {
	source := make([][]int, 0)
	remain := make([][]int, 0)
	if len(nums) < 2 {
		return nil, nil
	} else if len(nums) == 2 {
		return append(source, nums), nil
	} else if len(nums) == 3 {
		for i := 0; i < len(nums); i++ {
			s := make([]int, 0)
			r := make([]int, 0)
			r = append(r, nums[i])
			for j := 0; j < len(nums); j++ {
				if j != i {
					s = append(s, nums[j])
				}
			}
			source = append(source, s)
			remain = append(remain, r)
		}
	} else {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				s := make([]int, 0)
				r := make([]int, 0)
				r = append(r, nums[i])
				r = append(r, nums[j])
				for k := 0; k < len(nums); k++ {
					if k != i && k != j {
						s = append(s, nums[k])
					}
				}
				source = append(source, s)
				remain = append(remain, r)
			}
		}
	}
	return source, remain
}

func operator(nums []int) []int {
	if len(nums) != 2 {
		println("error")
		return nil
	} else {
		ans := make([]int, 4)
		var a int
		var b int
		if nums[0] > nums[1] {
			a = nums[0]
			b = nums[1]
		} else {
			a = nums[1]
			b = nums[0]
		}
		ans[0] = a + b
		ans[1] = a - b
		ans[2] = a * b
		ans[3] = b - a
		if b != 0 {
			ans = append(ans, a/b)
		}
		if a != 0 {
			ans = append(ans, b/a)
		}
		return ans
	}
}

func main() {
	//println(judgePoint24([]int{4,1,8,7}), true)
	println(judgePoint24([]int{1, 3, 4, 6}), true)
	println(judgePoint24([]int{8, 1, 6, 6}), true)
}
