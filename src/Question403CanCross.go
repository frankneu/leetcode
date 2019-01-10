package main

/**
https://leetcode.com/problems/frog-jump/
*/
func canCross(stones []int) bool {
	value := make([]map[int]int, 0)
	value = append(value, map[int]int{0: 0})
	for i := 1; i < len(stones); i++ {
		arr := make(map[int]int)
		for j := 0; j < i; j++ {
			step := stones[i] - stones[j]
			for _, v := range value[j] {
				if step == v-1 || step == v || step == v+1 {
					if _, ok := arr[step]; ok {
						continue
					} else {
						arr[step] = step
					}
				}
			}
		}
		value = append(value, arr)
	}
	return len(value[len(value)-1]) > 0
}

func main() {
	println(canCross([]int{0, 1, 3, 4, 5, 7, 9, 10, 12}))
}
