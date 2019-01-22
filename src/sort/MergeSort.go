package main

import "math"

func merge(arr []int, left, pivotal, right int) {
	nums1 := make([]int, pivotal-left+1)
	nums2 := make([]int, right-pivotal+1)
	copy(nums1, arr[left:pivotal])
	copy(nums2, arr[pivotal:right])
	nums1[len(nums1)-1] = math.MaxInt32
	nums2[len(nums2)-1] = math.MaxInt32
	i := 0
	j := 0
	for k := left; k < right; k++ {
		if nums1[i] < nums2[j] {
			arr[k] = nums1[i]
			i++
		} else {
			arr[k] = nums2[j]
			j++
		}
	}
}

func mergeSort(arr []int, left, right int) {
	if left < right {
		middle := (left + right) / 2
		mergeSort(arr, left, middle)
		mergeSort(arr, middle+1, right)
		merge(arr, left, middle, right)
	}
}

func main() {
	nums := []int{1, 2, 6, 7, 4, 3, 9, 8}
	mergeSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
