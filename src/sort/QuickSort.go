package main

import "math/rand"

/**
快速排序
*/
func quickSort(arr []int, left, right int) {
	if left < right {
		pivotal := partitionRandom(arr, left, right)
		quickSort(arr, left, pivotal-1)
		quickSort(arr, pivotal+1, right)
	}
}

/**
选择最后一个元素进行分区
*/
func partition(arr []int, left, right int) int {
	x := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

/**
随机选择划分点
*/
func partitionRandom(arr []int, left, right int) int {
	i := rand.Intn(right-left) + left
	arr[i], arr[right] = arr[right], arr[i]
	return partition(arr, left, right)
}

func main() {
	nums := []int{1, 2, 6, 7, 4, 3, 9, 8}
	quickSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
