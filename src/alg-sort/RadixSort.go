package main

import (
	"fmt"
	"strconv"
)

/**
将整数按位数切割成不同的数字，然后按每个位数分别比较。
具体做法是：
1. 将所有待比较数值统一为同样的数位长度，数位较短的数前面补零。
2. 然后，从最低位开始，依次进行一次排序。
3. 这样从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列。
*/

func RadixSort(arr []int) []int {
	if len(arr) < 2 {
		fmt.Println("NO NEED TO SORT")
		return arr
	}
	maxl := MaxLen(arr)
	return RadixCore(arr, 0, maxl)
}
func RadixCore(arr []int, digit, maxl int) []int { //核心排序机制时间复杂度 O( d( r+n ) )
	if digit >= maxl {
		return arr //排序稳定
	}
	radix := 10
	count := make([]int, radix)
	bucket := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		count[GetDigit(arr[i], digit)]++
	}
	for i := 1; i < radix; i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		d := GetDigit(arr[i], digit)
		bucket[count[d]-1] = arr[i]
		count[d]--
	}
	return RadixCore(bucket, digit+1, maxl)
}
func GetDigit(x, d int) int { //获取某位上的数字
	a := []int{1, 10, 100, 1000, 10000, 100000, 1000000}
	return (x / a[d]) % 10
}

func MaxLen(arr []int) int { //获取最大位数
	var maxl, curl int
	for i := 0; i < len(arr); i++ {
		curl = len(strconv.Itoa(arr[i]))
		if curl > maxl {
			maxl = curl
		}
	}
	return maxl
}

func main() {
	nums := []int{7, 6, 2, 1, 8, 5, 3}
	//nums2 := []int{1, 2, 6, 7, 4, 3, 9, 8}
	//merge(nums1, 0, 4, 6)
	nums = RadixSort(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
