package main

/**
https://leetcode.com/problems/4sum-ii/
*/
func fourSumCountV1(A []int, B []int, C []int, D []int) int {
	sumLeft := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			value := A[i] + B[j]
			if v, ok := sumLeft[value]; ok {
				sumLeft[value] = v + 1
			} else {
				sumLeft[value] = 1
			}
		}
	}

	sumRight := make(map[int]int)
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			value := C[i] + D[j]
			if v, ok := sumRight[value]; ok {
				sumRight[value] = v + 1
			} else {
				sumRight[value] = 1
			}
		}
	}

	result := 0
	for k, v := range sumLeft {
		if u, ok := sumRight[-k]; ok {
			result += u * v
		}
	}
	return result
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sumLeft := make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			sumLeft[A[i]+B[j]]++
		}
	}

	result := 0
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			result += sumLeft[-C[i]-D[j]]
		}
	}
	return result
}
