package main

import "math"

/**
左右都有序的数组进行合并
*/
func merge(arr []int, low, mid, high int) {
	leftLen := mid - low + 1
	rightLen := high - mid
	arrLeft := make([]int, leftLen+1)
	copy(arrLeft, arr[low:mid+1])
	arrLeft[leftLen] = math.MaxInt32 //哨兵牌
	arrRight := make([]int, rightLen+1)
	copy(arrRight, arr[mid+1:high+1])
	arrRight[rightLen] = math.MaxInt32 //哨兵牌
	i, j := 0, 0
	for k := low; k <= high; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}

/**
归并（Merge）排序法是将两个（或两个以上）有序表合并成一个新的有序表，即把待排序序列分为若干个子序列，每个子序列是有序的。然后再把有序子序列合并为整体有序序列。
*/
func mergeSort(arr []int, left, right int) {
	if left < right {
		middle := (left + right) / 2
		mergeSort(arr, left, middle)
		mergeSort(arr, middle+1, right)
		merge(arr, left, middle, right)
	}
}

func main() {
	nums := []int{7, 6, 2, 1, 8, 5, 3}
	//nums2 := []int{1, 2, 6, 7, 4, 3, 9, 8}
	//merge(nums1, 0, 4, 6)
	mergeSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
