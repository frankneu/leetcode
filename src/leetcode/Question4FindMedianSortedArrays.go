package main

func main() {
	var nums1 = []int{1, 2}
	var nums2 = []int{3, 4, 5, 6}
	println(findMedianSortedArrays(nums1, nums2))
}

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.
Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5

https://leetcode.com/problems/median-of-two-sorted-arrays/solution/
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}
	imin, imax, half_len := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := half_len - i
		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			max_of_left := 0
			min_of_right := 0
			if i == 0 {
				max_of_left = nums2[j-1]
			} else if j == 0 {
				max_of_left = nums1[i-1]
			} else {
				max_of_left = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(max_of_left)
			}
			if i == m {
				min_of_right = nums2[j]
			} else if j == n {
				min_of_right = nums1[i]
			} else {
				min_of_right = min(nums1[i], nums2[j])
			}
			return float64(max_of_left+min_of_right) / 2.0
		}
	}
	return 0.0
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

func findMedianSortedArraysV2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k := (m + n) / 2
	for k > 2 {
		if nums1[k/2] > nums2[k/2] {
			nums2 = nums2[k/2:]
		} else {
			nums1 = nums1[k/2:]
		}
		k = k - k/2
	}
	if k == 2 {
		return
	}
	return 0.0
}

func findK(nums1 []int, nums2 []int, k int, p int) float64 {
	m, n := len(nums1), len(nums2)

	if nums1[k/2] > nums2[k/2] {
		nums2 = nums2[k/2:]
	} else {
		nums1 = nums1[k/2:]
	}
	k = k - k/2

	for k > 2 {

	}
	if k == 2 {
		return
	}
	return 0.0
}
