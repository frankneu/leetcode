package main

import "math/rand"

/**
快速排序
1）选择一个基准元素,通常选择第一个元素或者最后一个元素,
2）通过一趟排序讲待排序的记录分割成独立的两部分，其中一部分记录的元素值均比基准元素值小。另一部分记录的 元素值比基准值大。
3）此时基准元素在其排好序后的正确位置
4）然后分别对这两部分记录用同样的方法继续进行排序，直到整个序列有序。
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
