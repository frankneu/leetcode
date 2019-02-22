package main

/**
https://leetcode.com/problems/intersection-of-two-arrays/
*/

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	result := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		set[nums1[i]] = nums1[i]
	}
	for i := 0; i < len(nums2); i++ {
		if v, ok := set[nums2[i]]; ok {
			if _, ok := result[v]; ok {
				continue
			} else {
				result[v] = v
			}
		}
	}
	ans := make([]int, len(result))
	i := 0
	for k := range result {
		ans[i] = k
		i++
	}
	return ans
}

func main() {
	//intersection([]int{1,2,2,1}, []int{2,2})
	v := intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	for i := 0; i < len(v); i++ {
		println(v[i])
	}
}
