package main

func main() {
	var nums1 = []int{1, 2, 3, 4, 5}
	var nums2 = []int{}
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
	start1, start2 := 0, 0
	end1 := len(nums1)
	end2 := len(nums2)
	value1, value2 := 0, 0
	var result float64 = 0.0
	for start1 < end1 || start2 < end2 {
		min1, min2, max1, max2 := 0, 0, 0, 0
		if start1 < len(nums1) {
			min1 = nums1[start1]
		}
		if end1 < len(nums1) {
			max1 = nums1[end1]
		}
		if start2 < len(nums1) {
			min2 = nums1[start2]
		}
		if end2 < len(nums1) {
			max2 = nums1[end2]
		}
		if min1 >= max2 {
			offset := (end1 - start1 + end2 - start2) / 2
			if offset > end2-start2 {
				value1 = nums1[start1+offset-end2+start2]
				value2 = nums1[start1+offset-end2+start2+1]
			} else if offset < end2-start2 {
				value1 = nums2[start2+offset]
				value2 = nums2[start2+offset+1]
			} else {
				value1 = nums2[end2]
				value2 = nums2[start2+offset+1]
			}
		}
		if min2 >= max1 {
			offset := (end1 - start1 + end2 - start2) / 2
			if offset > end1-start1 {
				value1 = nums2[start2+offset-end1+start1]
				value2 = nums2[start2+offset-end1+start1+1]
			} else {
				value1 = nums1[start1+offset-end2+start2]
				value2 = nums1[start1+offset-end1+start1+1]
			}
		}
	}
	length := len(nums1) + len(nums2)
	if length%2 == 1 {
		result = float64(value1+value2) / 2
	} else {
		result = float64(value1)
	}
	return result
}
