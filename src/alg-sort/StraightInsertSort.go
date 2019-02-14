package main

/**
将一个记录插入到已排序好的有序表中，从而得到一个新，记录数增1的有序表。
即：先将序列的第1个记录看成是一个有序的子序列;
然后从第2个记录逐个进行插入，直至整个序列有序为止。
要点：设立哨兵，作为临时存储和判断数组边界之用。
*/
func straightInsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		//0->i-1这段数据是已经排好序的，下面的循环作用就是找到这个序列里第一个比key小的元素
		for j > 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	nums := []int{1, 2, 6, 7, 4, 3, 9, 8}
	straightInsertSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
