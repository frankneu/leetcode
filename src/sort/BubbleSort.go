package main

//冒泡排序
/**
在要排序的一组数中，对当前还未排好序的范围内的全部数，自上而下对相邻的两个数依次进行比较和调整，让较大的数往下沉，较小的往上冒。
即：每当两相邻的数比较后发现它们的排序与排序要求相反时，就将它们互换。
*/
func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	nums := []int{1, 2, 6, 7, 4, 3, 9, 8}
	bubbleSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
