package main

/**
在要排序的一组数中，选出最小（或者最大）的一个数与第1个位置的数交换；
然后在剩下的数当中再找最小（或者最大）的与第2个位置的数交换;
依次类推，直到第n-1个元素（倒数第二个数）和第n个元素（最后一个数）比较为止。
*/
func simpleSelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		key := i
		for j := i + 1; j < len(arr); j++ {
			if arr[key] > arr[j] {
				key = j
			}
		}
		if key != i {
			arr[i], arr[key] = arr[key], arr[i]
		}
	}
}

func main() {
	nums := []int{1, 2, 6, 7, 4, 3, 9, 8}
	simpleSelectionSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
