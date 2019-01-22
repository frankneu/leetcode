package main

import "container/heap"

/**
https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
*/
type pair struct {
	a int
	b int
}

func (p *pair) value() int {
	return p.a + p.b
}

type PairHeap []pair

func (ph PairHeap) Len() int {
	return len(ph)
}

// 实现sort.Iterface
func (ph PairHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}
func (ph PairHeap) Less(i, j int) bool {
	return ph[i].value() < ph[j].value()
}

// 实现heap.Interface接口定义的额外方法
func (ph *PairHeap) Push(h interface{}) {
	*ph = append(*ph, h.(pair))
}
func (ph *PairHeap) Pop() (x interface{}) {
	n := len(*ph)
	x = (*ph)[n-1]    // 返回删除的元素
	*ph = (*ph)[:n-1] // [n:m]不包括下标为m的元素
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	result := make([][]int, 0)
	hp := &PairHeap{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			*hp = append(*hp, pair{nums1[i], nums2[j]})
		}
	}
	heap.Init(hp)
	for i := 0; i < k && hp.Len() > 0; i++ {
		p := heap.Pop(hp).(pair)
		result = append(result, []int{p.a, p.b})
	}
	return result
}

func main() {
	kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10)
}
