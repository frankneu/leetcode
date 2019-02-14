package main

/**
该方法的基本思想是：
先将整个待排元素序列分割成若干个子序列（由相隔某个“增量”的元素组成的）分别进行直接插入排序
然后依次缩减增量再进行排序
待整个序列中的元素基本有序（增量足够小）时,再对全体元素进行一次直接插入排序。
因为直接插入排序在元素基本有序的情况下（接近最好情况），效率是很高的，因此希尔排序在时间效率上比前两种方法有较大提高。
*/
func ShellSort(a []int) {
	n := len(a)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				swap(a, j, j-h)
			}
		}
		h /= 3
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	nums := []int{7, 6, 2, 1, 8, 5, 3}
	ShellSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
