package main

type NumArray struct {
	table [][]int
}

func Constructor(nums []int) NumArray {
	obj := NumArray{make([][]int, len(nums))}
	for i := 0; i < len(nums); i++ {
		obj.table[i] = make([]int, len(nums))
		for j := i; j < len(nums); j++ {
			if i == j {
				obj.table[i][j] = nums[i]
			} else {
				obj.table[i][j] = nums[j] + obj.table[i][j-1]
			}
		}
	}
	return obj
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.table[i][j]
}

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	println(obj.SumRange(0, 2))
}
